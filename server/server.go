package server

import (
	// Standard lib
	"time"

	// Internal
	"github.com/marksost/img/config"
	"github.com/marksost/img/helpers"

	// Third-party
	log "github.com/Sirupsen/logrus"
	"github.com/kataras/iris"
)

const (
	// URL Param used to indicate a debug request
	DEBUG_PARAM = "debug"
	// Key to store response headers under in the request context
	RESPONSE_HEADERS_KEY = "response-headers"
)

var (
	// Server for all requests
	server *iris.Framework
)

func Start() {
	// Get configuration instance
	c := config.GetInstance()

	// Set up middleware
	setMiddleWare()

	// Set up custom serializers
	setSerializers()

	// Set up routes
	setRoutes()

	// Set up server configuration
	server.Config.Gzip = true
	server.Config.ReadTimeout = time.Duration(c.Server.Timeouts.Read) * time.Second
	server.Config.WriteTimeout = time.Duration(c.Server.Timeouts.Write) * time.Second

	// Attempt to start the server
	go server.Listen(":" + helpers.Int2String(c.Server.Port))
}

func Stop() {
	// Catch panics
	// NOTE: These can occur when the network connection is already closed
	defer func() {
		_ = recover()
	}()

	// Attempt to stop the server
	if err := server.Close(); err != nil {
		log.WithField("error", err.Error()).Warn("Error stopping the server")
	}
}

func init() {
	// Create new server instance
	server = iris.New()
}
