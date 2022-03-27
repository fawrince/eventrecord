package main

import (
	"flag"
	"github.com/fawrince/eventrecord/application"
	"github.com/fawrince/eventrecord/http"
	"github.com/fawrince/eventrecord/logger"
	"os"
	"os/signal"
	"syscall"
)

var sigusr1 = make(chan os.Signal, 1)
var sigterm = make(chan os.Signal, 1)

func init() {
	signal.Notify(sigusr1, syscall.SIGUSR1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	flag.Parse()

	/*file, err := os.OpenFile("all.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open out.log")
	}
	defer file.Close()*/

	logger := logger.NewLogger(os.Stdout)

	app := application.NewApp(logger)
	app.Setup(sigterm, sigusr1)

	srv := http.NewServer(logger, app)
	srv.Start()

	<-sigterm
	srv.Stop()
	app.Stop()
}
