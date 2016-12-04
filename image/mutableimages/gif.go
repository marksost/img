// gif contains all functionality around doing the actual image processing actions
// that are supported by the application on GIF images
package mutableimages

import (
	// Standard lib
	"bytes"
	"fmt"
	"image/gif"
	"os/exec"

	// Internal
	"github.com/marksost/img/values"
)

const (
	// The command to exec when manipulating a GIF
	GIF_COMMAND = "gifsicle"
	// The command argument used to crop a GIF
	GIF_CROP_COMMAND = "--crop=%d,%d+%dx%d"
	// The command argument used to resize a GIF
	GIF_RESIZE_COMMAND = "--resize=%dx%d"
)

type (
	// Struct representing a process-able GIF image
	GifMutableImage struct {
		decodedData *gif.GIF          // The decoded data from the image
		img         *ProcessableImage // The processable image struct containing all image information
		width       int               // The current width of the image
		height      int               // The current height of the image
	}
)

// NewGifMutableImage creates a new `GifMutableImage` and returns it
func NewGifMutableImage(img *ProcessableImage) (*GifMutableImage, error) {
	// Set error for use in this method
	var err error

	// Form new image
	i := &GifMutableImage{
		img: img,
	}

	// Attempt to decode the image
	i.decodedData, err = gif.DecodeAll(bytes.NewBuffer(img.Data))
	if err != nil {
		return nil, err
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
func (i *GifMutableImage) GetWidth() int64 {
	return int64(i.width)
}

// GetWidth returns the width of the image
// NOTE: This width is based on the current image data available,
// meaning it may be different from the original width of the image
// if resize operations have occurred
func (i *GifMutableImage) GetHeight() int64 {
	return int64(i.height)
}

/* End dimension methods */

/* Begin operation methods */

// Crop performs a crop operation on the image
// based on input width/height/x/y values
func (i *GifMutableImage) Crop(vals *values.CropValues) error {
	// Form command arguments
	args := []string{
		fmt.Sprintf(GIF_CROP_COMMAND, vals.X, vals.Y, vals.Width, vals.Height),
	}

	// Return value of internal command call
	return i.runCommand(args)
}

// Resize performs a resize operation on the image
// based on input width/height values
func (i *GifMutableImage) Resize(vals *values.DimensionValues) error {
	// Form command arguments
	args := []string{
		fmt.Sprintf(GIF_RESIZE_COMMAND, vals.Width, vals.Height),
	}

	// Return value of internal command call
	return i.runCommand(args)
}

/* End operation methods */

/* Begin internal property methods */

// Img returns a mutable image's processable image property
func (i *GifMutableImage) Img() *ProcessableImage {
	return i.img
}

// SetDefault is used to set any needed default values for the mutable image
// before image processing starts
func (i *GifMutableImage) SetDefaults() {}

// SetDimensions reads in an image and sets it's dimensions
func (i *GifMutableImage) SetDimensions() {
	// Read bounds from image data
	bounds := i.decodedData.Image[0].Bounds()

	// Set width and height
	i.width = bounds.Dx()
	i.height = bounds.Dy()
}

/* End internal property methods */

/* Begin utility methods */

// runCommand runs a GIF command on the host system, with one or more arguments passed in
// and returns the data returned by the command when possible
func (i *GifMutableImage) runCommand(args []string) error {
	var (
		// Set error for use in this method
		err error
		// Command output
		output bytes.Buffer
	)

	// Form command
	cmd := exec.Command(GIF_COMMAND, args...)
	cmd.Stdin = bytes.NewReader(i.img.Data)
	cmd.Stdout = &output

	// Run command
	if err = cmd.Run(); err != nil {
		return err
	}

	// Reset image data and attempt to decode the image
	i.img.Data = output.Bytes()
	if i.decodedData, err = gif.DecodeAll(bytes.NewBuffer(i.img.Data)); err != nil {
		return err
	}

	// Reset dimensions for the image data
	i.SetDimensions()

	return nil
}

/* End utility methods */
