package server

import (
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
	resp := OKResponse
	resp.Data = map[string]string{
		"img": c.Param("img"),
	}

	// Write JSON output
	c.JSON(resp.Code, resp)
}

// Handles all dis-allowed routes
func methodNotAllowed(c *iris.Context) {
	// Write JSON output
	c.JSON(MethodNotAllowedResponse.Code, MethodNotAllowedResponse)
}

// Handles all 404 server errors
func notFound(c *iris.Context) {
	// Write JSON output
	c.JSON(NotFoundResponse.Code, NotFoundResponse)
}

// Handles all routes server 200 OK statuses
func ok(c *iris.Context) {
	// Write JSON output
	c.JSON(OKResponse.Code, nil)
}

// Handles all 500 server errors
func serverError(c *iris.Context) {
	// Write JSON output
	c.JSON(ServerErrorResponse.Code, ServerErrorResponse)
}
