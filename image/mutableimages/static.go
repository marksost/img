// static contains all functionality around doing the actual image processing actions
// that are supported by the application on "static" images
// NOTE: "static" images are any non-animated (GIF) images
package mutableimages

import (
	// Internal
	"github.com/marksost/img/config"

	// Third-party
	"github.com/h2non/bimg"
)

type (
	// Struct representing a process-able "static" image
	StaticMutableImage struct {
		img     *ProcessableImage // The processable image struct containing all image information
		quality int               // The quality to use when outputing the processed image
		width   int               // The current width of the image
		height  int               // The current height of the image
	}
)

// NewStaticMutableImage creates a new `StaticMutableImage` and returns it
func NewStaticMutableImage(img *ProcessableImage) (*StaticMutableImage, error) {
	// Form new image
	i := &StaticMutableImage{
		img: img,
	}

	// Set dimensions for the image data
	i.SetDimensions()

	return i, nil
}

// GetWidth returns the width of the image
// NOTE: This width is based on the current image data available,
// meaning it may be different from the original width of the image
// if resize operations have occurred
func (i *StaticMutableImage) GetWidth() int64 {
	return int64(i.width)
}

// GetWidth returns the width of the image
// NOTE: This width is based on the current image data available,
// meaning it may be different from the original width of the image
// if resize operations have occurred
func (i *StaticMutableImage) GetHeight() int64 {
	return int64(i.height)
}

// Img returns a mutable image's processable image property
func (i *StaticMutableImage) Img() *ProcessableImage {
	return i.img
}

// SetDefault is used to set any needed default values for the mutable image
// before image processing starts
func (i *StaticMutableImage) SetDefaults() {
	// Set default quality
	i.quality = config.GetInstance().Images.DefaultQuality
}

// SetDimensions reads in an image and sets it's dimensions
func (i *StaticMutableImage) SetDimensions() {
	// Read size from image data
	size, err := bimg.NewImage(i.img.Data).Size()
	if err != nil {
		return
	}

	// Set width and height
	i.width = size.Width
	i.height = size.Height
}
