package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/golangbg/site/twitterbot"
	"github.com/golangbg/site/webserver"
)

func main() {
	// Create a logger
	log := logrus.New()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Start the webserver
	log.Infoln("[webserver] starting")
	srv := webserver.New(log)
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Infoln("[webserver] started")

	// Start the twitterbot
	log.Infoln("[twitterbot] starting")
	tb, err := twitterbot.New(log, "#golangbg", "#golang")
	if err != nil {
		log.Errorf("[twitterbot] %v", err)
	} else {
		go tb.Start()
		log.Infoln("[twitterbot] started")
	}

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Infoln("received SIGINT")
	case syscall.SIGTERM:
		log.Infoln("received SIGTERM")
	}

	log.Infoln("shutting down...")
	srv.Shutdown(context.Background())
	tb.Shutdown()
	log.Infoln("shutdown completed")
}
