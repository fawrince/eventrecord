package application

import (
	"github.com/fawrince/eventrecord/broker"
	"github.com/fawrince/eventrecord/dto"
	"github.com/fawrince/eventrecord/logger"
	"os"
	"time"
)

type ProduceHandler interface {
	ProduceInput() chan<- dto.Coordinates
	Stop()
}

type ConsumeHandler interface {
	ConsumeOutput() <-chan dto.Coordinates
	Stop()
}

type App struct {
	logger   *logger.Logger
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

	app.Produce = producer
	app.Consume = consumer
}

func (app *App) Stop() {
	logger := app.logger

	app.Produce.Stop()
	app.Consume.Stop()

	logger.Infof("Application stopped")
}
