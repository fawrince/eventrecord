package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fawrince/eventrecord/dto"
	"github.com/fawrince/eventrecord/logger"
	"log"
	"os"
	"strings"
	"sync"
)

type Consumer struct {
	client sarama.ConsumerGroup
	ready  chan bool
	close  chan bool
	logger *logger.Logger
	output chan dto.Coordinates
}

func NewConsumer(logger *logger.Logger) *Consumer {
	return &Consumer{
		logger: logger,
		ready:  make(chan bool),
		close:  make(chan bool),
		output: make(chan dto.Coordinates, 100),
	}
}

// ConsumeOutput returns the pipelined-channel to consume coordinates
func (consumer *Consumer) ConsumeOutput() <-chan dto.Coordinates {
	return consumer.output
}

func (consumer *Consumer) Start(sigterm chan os.Signal, sigusr1 chan os.Signal) {
	consumer.logger.Infof("Starting a new Sarama consumer at %s", Brokers)

	if Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(Version)
	if err != nil {
		consumer.logger.Fatal(fmt.Errorf("error parsing Kafka version: %w", err))
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(Brokers, ","), Group, config)
	if err != nil {
		consumer.logger.Fatal(fmt.Errorf("error creating consumer group client: %w", err))
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(Topics, ","), consumer); err != nil {
				consumer.logger.Fatal(fmt.Errorf("error from consumer: %w", err))
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	consumer.logger.Infof("Consumer is up and running!...")
	consumer.client = client

	/*
		consumptionIsPaused := false
		for keepRunning {
			select {
			case <-ctx.Done():
				consumer.logger.Infof("terminating consumer: context cancelled")
				keepRunning = false
			case <-sigterm:
				consumer.logger.Infof("terminating consumer: via signal")
				keepRunning = false
			case <-sigusr1:
				toggleConsumptionFlow(client, &consumptionIsPaused)
			}
		}*/

	go func() {
		select {
		case <-consumer.close:
			cancel()
			wg.Wait()
			consumer.logger.Infof("terminating consumer: close call")
		case <-ctx.Done():
			consumer.logger.Infof("terminating consumer: context cancelled")
		case <-sigterm:
			consumer.logger.Infof("terminating consumer: via signal")
		}
	}()
}

func (consumer *Consumer) Stop() {
	consumer.close <- true

	if err := consumer.client.Close(); err != nil {
		consumer.logger.Fatal(fmt.Errorf("error closing client: %w", err))
	}
	consumer.logger.Infof("Consumer is successfully closed")
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		var coord dto.Coordinates
		json.Unmarshal(message.Value, &coord)
		consumer.logger.Infof("Message claimed: value = %v, timestamp = %v, topic = %s", coord, message.Timestamp, message.Topic)
		consumer.output <- coord
		session.MarkMessage(message, "")
	}
	return nil
}
