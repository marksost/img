package server

import (
	// Standard lib
	"net/http"

	// Internal
	"github.com/marksost/img/image"

	// Third-party
	"github.com/kataras/iris"
)

func setRoutes() {
	// Main image handling route
	server.Get("/*img", img)

	// HEAD requests
	server.Head("/*img", ok)

	// Dis-allowed routes/methods
	server.Post("/*img", methodNotAllowed)
	server.Put("/*img", methodNotAllowed)
	server.Patch("/*img", methodNotAllowed)
	server.Delete("/*img", methodNotAllowed)

	// Handle errors
	server.OnError(NotFoundResponse.Code, notFound)
	server.OnError(ServerErrorResponse.Code, serverError)
}

// Handles all GEt requests to the application
// not matching any other route rules
func img(c *iris.Context) {
	// Form new image
	i := &image.Image{}

	// Set default response
	var resp *Response = OKResponse

	// Process request
	if err := i.Process(c); err != nil {
		// Store error code
		code := err.(*image.ImageRequestError).Code()

		resp = &Response{
			Code:    code,
			Message: http.StatusText(code),
			Data:    []string{err.Error()},
		}
	}

	// Write JSON output
	JSON(c, resp)
}

// Handles all dis-allowed routes
func methodNotAllowed(c *iris.Context) {
	// Write JSON output
	JSON(c, MethodNotAllowedResponse)
}

// Handles all 404 server errors
func notFound(c *iris.Context) {
	// Write JSON output
	JSON(c, NotFoundResponse)
}

// Handles all routes server 200 OK statuses
func ok(c *iris.Context) {
	// Write JSON output
	JSON(c, OKResponse)
}

// Handles all 500 server errors
func serverError(c *iris.Context) {
	// Write JSON output
	JSON(c, ServerErrorResponse)
}
