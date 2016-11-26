// mutableimages contains all functionality around doing the actual image processing actions
// that are supported by the application
package mutableimages

import (
	// Standard lib
	"fmt"

	// Internal
	"github.com/marksost/img/image/utils"
)

type (
	// Interface describing methods used to process an image
	MutableImage interface {
		// Dimension methods
		GetWidth() int64
		GetHeight() int64

		// Operation methods
		// TO-DO: Write these and their signatures
		Resize(int64, int64) error

		// Internal property methods
		Img() *ProcessableImage
		SetDefaults()
		SetDimensions()
	}
	// Struct representing a set of dta to be used when processing mutable images
	ProcessableImage struct {
		Animated     bool   // Whether the image is animated or not
		Data         []byte // Image data to be used for processing
		ImageType    string // The MIME type of the image
		SourceWidth  int64  // The initial width of the image
		SourceHeight int64  // The initial height of the image
	}
)

// NewMutableImage creates a new `MutableImage` and returns it
func NewMutableImage(data []byte, imageType string) (MutableImage, error) {
	var (
		// Set error for use in this method
		err error
		// Set default mutable image
		mi MutableImage
		// Create processable image
		pi *ProcessableImage = &ProcessableImage{
			Animated:  imageType == utils.GIF_MIME,
			Data:      data,
			ImageType: imageType,
		}
	)

	// Generate mutable image based on image type
	switch imageType {
	case utils.GIF_MIME:
		// Create GIF mutable image
		mi, err = NewGifMutableImage(pi)
		if err != nil {
			return nil, err
		}
	case utils.JPEG_MIME, utils.PNG_MIME, utils.TIFF_MIME:
		// Create static mutable image
		mi, err = NewStaticMutableImage(pi)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported image type: %s", imageType)
	}

	// Set default values for mutable image
	mi.SetDefaults()

	// Set dimensions for processable image
	pi.SourceWidth = mi.GetWidth()
	pi.SourceHeight = mi.GetHeight()

	return mi, nil
}
