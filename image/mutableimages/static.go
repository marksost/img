// static contains all functionality around doing the actual image processing actions
// that are supported by the application on "static" images
// NOTE: "static" images are any non-animated (GIF) images
package mutableimages

import (
	// Internal
	"github.com/marksost/img/config"
	"github.com/marksost/img/values"

	// Third-party
	"github.com/h2non/bimg"
)

type (
	// Struct representing a process-able "static" image
	StaticMutableImage struct {
		// The processable image struct containing all image information
		img *ProcessableImage
		// Max-width of the image before switching interpolators
		interpolatorThreshold int64
		// The current width of the image
		width int
		// The current height of the image
		height int
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

/* Begin dimension methods */

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

/* End dimension methods */

/* Begin operation methods */

// Crop performs a crop operation on the image
// based on input width/height/x/y values
func (i *StaticMutableImage) Crop(vals *values.CropValues) error {
	// Form options
	opts := bimg.Options{
		Top:        int(vals.Y),
		Left:       int(vals.X),
		AreaWidth:  int(vals.Width),
		AreaHeight: int(vals.Height),
		Quality:    100,
	}

	// Reset top if the point is the top-left corner
	if opts.Top == 0 && opts.Left == 0 {
		opts.Top = -1
	}

	// Return value of internal resize call
	return i.resize(opts)
}

// Quality performs a quality operation on the image
// based on input value
func (i *StaticMutableImage) Quality(val int64) error {
	// Form options
	opts := bimg.Options{
		Interlace:      true,
		Interpretation: bimg.InterpretationSRGB,
		Quality:        int(val),
	}

	// Return value of internal resize call
	return i.resize(opts)
}

// Resize performs a resize operation on the image
// based on input width/height values
func (i *StaticMutableImage) Resize(vals *values.DimensionValues) error {
	// Form options
	opts := bimg.Options{
		Width:        int(vals.Width),
		Height:       int(vals.Height),
		Quality:      100,
		Force:        true,
		Interpolator: bimg.Bilinear,
	}

	// Switch interpolator if needed
	if vals.Width <= i.interpolatorThreshold {
		opts.Interpolator = bimg.Bicubic
	}

	// Return value of internal resize call
	return i.resize(opts)
}

/* End operation methods */

/* Begin internal property methods */

// Img returns a mutable image's processable image property
func (i *StaticMutableImage) Img() *ProcessableImage {
	return i.img
}

// SetDefault is used to set any needed default values for the mutable image
// before image processing starts
func (i *StaticMutableImage) SetDefaults() {
	// Set defaults
	i.interpolatorThreshold = int64(config.GetInstance().Images.InterpolatorThreshold)
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

/* End internal property methods */

/* Begin utility methods */

// resize takes a set of bimg options and calls for a `resize` on the image data
// NOTE: `resize` handles more than just resizing of an image (ex: cropping)
func (i *StaticMutableImage) resize(opts bimg.Options) error {
	// Resize image with opts
	data, err := bimg.Resize(i.img.Data, opts)
	if err != nil {
		return err
	}

	// Reset image data
	i.img.Data = data

	// Reset dimensions for the image data
	i.SetDimensions()

	return nil
}

/* End utility methods */
