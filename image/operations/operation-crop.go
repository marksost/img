package operations

import (
	// Standard lib
	"fmt"
	"strings"

	// Internal
	"github.com/marksost/img/helpers"
	"github.com/marksost/img/image/mutableimages"
	"github.com/marksost/img/values"
)

type (
	// Struct representing a crop operation to be performed on an image
	CropOperation struct {
		// Mutable image to use when processing this operation
		mi mutableimages.MutableImage
		// Raw query string value for this operation
		rawValue string
		// Values used when operating on the image
		values *values.CropValues
	}
)

// Process is used to perform the actual operation processing
// on a given image
func (o *CropOperation) Process(mi *mutableimages.MutableImage) error {
	// Set internal value
	o.mi = *mi
	o.values = &values.CropValues{}

	// Parse raw value
	if err := o.parse(); err != nil {
		return err
	}

	// Validate operation
	if err := o.Validate(); err != nil {
		return err
	}

	// Return value from crop operation
	return o.mi.Crop(o.values)
}

// String returns a string representation of this operation
func (o *CropOperation) String() string {
	// Validate operation
	if err := o.Validate(); err != nil {
		return ""
	}

	// Form return value
	str := OPERATION_NAME_CROP + QUERY_STRING_ENTRY_DELIMITER
	str += "{w}" + values.DIMENSION_DELIMITER + "{h}"
	str += values.CROP_DELIMITER + "{x}" + values.POINT_DELIMITER + "{y}"

	// Replace macros
	str = strings.Replace(str, "{w}", helpers.Int642String(o.values.Width), -1)
	str = strings.Replace(str, "{h}", helpers.Int642String(o.values.Height), -1)
	str = strings.Replace(str, "{x}", helpers.Int642String(o.values.X), -1)
	str = strings.Replace(str, "{y}", helpers.Int642String(o.values.Y), -1)

	return str
}

// Validate returns a boolean indicating if the operation can be run,
// including checking source image against proposed operation parameters
func (o *CropOperation) Validate() error {
	// Verify values exist
	if o.values == nil {
		return fmt.Errorf("Invalid values. Operation appears to not have been initialized")
	}

	// Verify target values
	if o.values.Width == 0 || o.values.Height == 0 {
		return fmt.Errorf("Invalid target values detected: %v", o.values)
	}

	// Verify operation dimensions are within image bounds
	if o.values.X+o.values.Width > o.mi.GetWidth() ||
		o.values.Y+o.values.Height > o.mi.GetHeight() {
		return fmt.Errorf("Target values exceed image bounds")
	}

	return nil
}

// parse is used to parse an operation's raw value and convert it
// into usable data for the operation
func (o *CropOperation) parse() error {
	// Split raw value based on crop delimiter and validate result
	bits := strings.Split(o.rawValue, values.CROP_DELIMITER)
	if len(bits) != 2 {
		return fmt.Errorf("Invalid values passed in for operation")
	}

	// Split dimension and point values and validate result
	db := strings.Split(bits[0], values.DIMENSION_DELIMITER)
	pb := strings.Split(bits[1], values.POINT_DELIMITER)
	if len(db) != 2 || len(pb) != 2 {
		return fmt.Errorf("Invalid values passed in for operation")
	}

	// Get dimension values
	dv, err := values.NewDimensionValues(db[0], db[1], o.mi.GetWidth(), o.mi.GetHeight())
	if err != nil {
		return err
	}

	// Get point values
	pv, err := values.NewPointValues(pb[1], pb[0], o.mi.GetWidth(), o.mi.GetHeight())
	if err != nil {
		return err
	}

	// Set values
	o.values = &values.CropValues{
		Width:  dv.Width,
		Height: dv.Height,
		X:      pv.X,
		Y:      pv.Y,
	}

	return nil
}
