package main

import (
	// Standard lib
	"flag"
	"os"
	"os/signal"
	"runtime"

	// Internal
	"github.com/marksost/img/config"

	// Third-party
	log "github.com/Sirupsen/logrus"
)

func main() {
	// Log start of the service
	log.Info("Application is starting")

	// Set max processes
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Initialize configuration
	config.Init()

	// Get configuration instance
	c := config.GetInstance()

	// Log configuration value only in development environments
	if c.IsDevelopment() {
		log.WithField("config", c).Info("Configuration")
	}

	// Parse flags
	flag.Parse()

	// Listen for and exit the application on SIGKILL or SIGINT
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, os.Kill)

	select {
	case <-stop:
		log.Info("Server is shutting down")
	}
}
