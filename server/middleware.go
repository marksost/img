package server

import (
	// Internal
	"github.com/marksost/img/config"

	// Third-party
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
)

// setMiddleWare is used to set up middleware for all requests
func setMiddleWare() {
	// Set up logging middleware
	server.Use(logger.New())

	// Set up CORS
	server.Use(cors.Default())

	// Handle recovery from panics
	server.Use(recovery.New())

	// Add pre-processing response headers
	server.UseFunc(addPreflightResponseHeaders)

	// Add post-processing response headers
	server.DoneFunc(addPostflightResponseHeaders)
}

// addPreflightResponseHeaders is used to add common response headers to requests
// NOTE: These headers are added *before* processing the request
func addPreflightResponseHeaders(c *iris.Context) {
	// Add application name header
	c.SetHeader("X-Powered-By", config.GetInstance().Name)

	// Go to next middleware
	c.Next()
}

// addPostflightResponseHeaders is used to add response headers to requests
// NOTE: These headers are added *after* processing the request
func addPostflightResponseHeaders(c *iris.Context) {
	// TO-DO: Add post-processing headers if needed

	// Go to next middleware
	c.Next()
}
