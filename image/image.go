package image

import (
	// Third-party
	"github.com/kataras/iris"
)

const (
	// The default MIME type to return when no MIME type was set from the downloaded image
	DEFAULT_MIME_TYPE = "application/octet-stream"
)

type (
	// Struct representing a single image to be processed from a HTTP request
	Image struct {
		ctx        *iris.Context // The request context this image relates to
		mimeType   string        // The detected MIME type of the downloaded image
		outputData []byte        // The processed image as data
	}
)

/* Begin main public functionality methods */

// NewImage creates a new `Image` and returns it
func NewImage(ctx *iris.Context) *Image {
	// Create and return new image with context set from input
	return &Image{
		ctx:        ctx,
		outputData: make([]byte, 0),
	}
}

// Process is used to handle a single image processing request
// It will download the image, process it based on request parameters
// and return the result
func (i *Image) Process() error {
	// TO-DO: Fill out download and process functionality

	return NewError(405, "This is a test")
}

/* End main public functionality methods */

/* Begin internal propery methods */

// Data returns a byte slice representing the processed image
func (i *Image) Data() []byte {
	return i.outputData
}

// MimeType returns a string representing the MIME type of the downloaded image
// NOTE: will return a default MIME type if none was previously set
func (i *Image) MimeType() string {
	// Check for empty MIME type and return default
	if i.mimeType == "" {
		return DEFAULT_MIME_TYPE
	}

	return i.mimeType
}

/* End internal propery methods */
