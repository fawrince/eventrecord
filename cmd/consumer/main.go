package main

import (
	"github.com/fawrince/eventrecord/broker"
	"github.com/fawrince/eventrecord/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	logger := logger.NewLogger(os.Stdout)

	time.Sleep(time.Second * 15)

	logger.Infof("Try start application")

	consumer := broker.NewConsumer(logger)
	consumer.Start(sigterm, sigusr1)
}
