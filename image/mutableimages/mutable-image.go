package mutableimages

import (
	// Standard lib
	"fmt"

	// Internal
	"github.com/marksost/img/image/utils"
)

type (
	MutableImage interface {
		// Dimension methods
		GetWidth() int64
		GetHeight() int64

		// Operation methods
		// TO-DO: Write these and their signatures
		// Crop()
		// Density()
		// Resize()
		// Quality()

		// Internal property methods
		SetDefaults()
	}
	ProcessableImage struct {
		Animated     bool
		Data         []byte
		DataSize     int64
		ImageType    string
		Width        int64
		Height       int64
		SourceWidth  int64
		SourceHeight int64
	}
)

// NewMutableImage creates a new `MutableImage` and returns it
func NewMutableImage(data []byte, imageType string) (MutableImage, error) {
	// Create mutable image and processable image
	var (
		// Set error for use in this method
		err error
		// Set default mutable image
		mi MutableImage
		// Create processable image
		pi *ProcessableImage = &ProcessableImage{
			Animated:  imageType == utils.GIF_MIME,
			Data:      data,
			DataSize:  int64(len(data)),
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
	pi.Width = mi.GetWidth()
	pi.Height = mi.GetHeight()
	pi.SourceWidth = pi.Width
	pi.SourceHeight = pi.Height

	return mi, nil
}
