package broker

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/fawrince/eventrecord/logger"
)

type Producer struct {
	logger   *logger.Logger
	producer sarama.AsyncProducer
	address  string
	topic    string
}

func NewProducer(logger *logger.Logger, address string, topic string) *Producer {
	return &Producer{
		logger:   logger,
		producer: nil,
		address:  address,
		topic:    topic,
	}
}

// ProduceInput returns the pipelined-channel to produce coordinates
func (prod *Producer) ProduceInput() chan<- Coordinates {
	ch := make(chan Coordinates, 100)
	go func() {
		for coordinates := range ch {
			data, _ := json.Marshal(coordinates)
			message := &sarama.ProducerMessage{
				Topic: prod.topic,
				Key:   sarama.StringEncoder(coordinates.Client),
				Value: sarama.ByteEncoder(data),
			}
			prod.producer.Input() <- message
			prod.logger.Tracef("Message produced: %v", coordinates)
		}
	}()
	return ch
}

func (prod *Producer) Start() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	prod.logger.Infof("Try start the producer at %s", prod.address)
	producer, err := sarama.NewAsyncProducer([]string{prod.address}, config)
	if err != nil {
		panic(err)
	}
	prod.producer = producer
	prod.logger.Infof("Producer started")

	var (
		successes, errors int
	)

	go func() {
		for range prod.producer.Successes() {
			successes++
		}
		prod.logger.Infof("Producer stopped with %d successes", successes)
	}()

	go func() {
		for err := range prod.producer.Errors() {
			prod.logger.Error(err)
			errors++
		}
		prod.logger.Infof("Producer stopped with %d errors", errors)
	}()
}

func (prod *Producer) Stop() {
	prod.producer.AsyncClose()
	prod.logger.Infof("Producer stopped asynchronously")
}
