package image

import (
	// Third-party
	"github.com/kataras/iris"
)

type (
	// Struct representing a single image to be processed from a HTTP request
	Image struct {
		ctx *iris.Context // The request context this image relates to
	}
)

// Process is used to handle a single image processing request
// It will download the resource, process it based on request parameters
// and return the result
func (i *Image) Process(ctx *iris.Context) error {
	// Set request context
	i.ctx = ctx

	// TO-DO: Fill out download and process functionality

	return NewError(405, "This is a test")
}
