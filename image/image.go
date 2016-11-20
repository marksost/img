package image

import (
	// Standard lib
	"net/http"
	"net/url"

	// Internal
	"github.com/marksost/img/image/utils"

	// Third-party
	"github.com/kataras/iris"
)

type (
	// Struct representing a single image to be processed from a HTTP request
	Image struct {
		ctx        *iris.Context // The request context this image relates to
		outputData []byte        // The processed image as data
		utils      *ImageUtils   // A collection of utilities used while processing a request
	}
	// Struct representing an `Image` struct's utilities used while processing a request
	ImageUtils struct {
		Downloader *utils.Downloader // Utility used to form URLs and download images from them
	}
)

/* Begin main public functionality methods */

// NewImage creates a new `Image` and returns it
func NewImage(ctx *iris.Context) *Image {
	// Create and return new image with context set from input
	return &Image{
		ctx:        ctx,
		outputData: make([]byte, 0),
		utils: &ImageUtils{
			Downloader: utils.NewDownloader(ctx.Param("img")),
		},
	}
}

// Process is used to handle a single image processing request
// It will download the image, process it based on request parameters
// and return the result
func (i *Image) Process() error {
	// Use downloader utility to download image from URL
	err := i.utils.Downloader.Download()
	if err != nil {
		// Return bad request error
		return NewError(http.StatusBadRequest, err.Error())
	}

	// TO-DO: Fill out process functionality

	return NewError(http.StatusOK, "This is a test")
}

/* End main public functionality methods */

/* Begin internal propery methods */

// Data returns a byte slice representing the processed image
func (i *Image) Data() []byte {
	return i.outputData
}

/* End internal propery methods */

/*  Begin utils proxy methods */

// MimeType returns a string representing the MIME type of the downloaded image
// NOTE: will return a default MIME type if none was previously set
// NOTE: Proxies the call to this image's downloader utility
func (i *Image) MimeType() string {
	return i.utils.Downloader.MimeType()
}

// RawData returns a byte slice representing the raw data from the downloaded image
// NOTE: Proxies the call to this image's downloader utility
func (i *Image) RawData() []byte {
	return i.utils.Downloader.Data()
}

// Url returns a URL struct representing the parsed URL of the requested image
// NOTE: Proxies the call to this image's downloader utility
func (i *Image) Url() *url.URL {
	return i.utils.Downloader.Url()
}

/*  End utils proxy methods */

/* Begin utility methods */
/* End utility methods */
