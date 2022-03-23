package application

import (
	"github.com/fawrince/eventrecord/broker"
	"github.com/fawrince/eventrecord/logger"
	"os"
	"time"
)

type ProduceHandler interface {
	ProduceInput() chan<- broker.Coordinates
}

type ConsumeHandler interface {
	ConsumeOutput() <-chan broker.Coordinates
}

type App struct {
	logger   *logger.Logger
	producer *broker.Producer
	consumer *broker.Consumer
	Produce  ProduceHandler
	Consume  ConsumeHandler
}

func NewApp(logger *logger.Logger) *App {
	return &App{
		logger: logger,
	}
}

func (app *App) Setup(sigterm chan os.Signal, sigusr1 chan os.Signal) {
	logger := app.logger

	logger.Infof("Wait 30 second to the broker startup...")
	time.Sleep(time.Second * 30)

	logger.Infof("Setup the producer...")
	producer := broker.NewProducer(logger, broker.Brokers, broker.Topics)
	producer.Start()

	logger.Infof("Setup the consumer...")
	consumer := broker.NewConsumer(logger)
	consumer.Start(sigterm, sigusr1)

	app.producer = producer
	app.consumer = consumer

	app.Produce = producer
	app.Consume = consumer
}

func (app *App) Stop() {
	logger := app.logger

	app.producer.Stop()
	app.consumer.Stop()

	logger.Infof("Application stopped")
}
