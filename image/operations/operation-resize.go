package operations

import (
	// Standard lib
	"fmt"
	"strings"

	// Internal
	"github.com/marksost/img/helpers"
	"github.com/marksost/img/image/mutableimages"
)

type (
	// Struct representing a resize operation to be performed on an image
	ResizeOperation struct {
		// Struct representing target dimensions for the image
		dimensions *DimensionValues
		// Mutable image to use when processing this operation
		mi mutableimages.MutableImage
		// Raw query string value for this operation
		rawValue string
	}
)

// Process is used to perform the actual operation processing
// on a given image
func (o *ResizeOperation) Process(mi *mutableimages.MutableImage) error {
	// Set internal value
	o.mi = *mi

	// Parse raw value
	if err := o.parse(); err != nil {
		return err
	}

	// Validate operation
	if err := o.Validate(); err != nil {
		return err
	}

	// Return value from resize operation
	return o.mi.Resize(o.dimensions.Width, o.dimensions.Height)
}

// String returns a string representation of this operation
func (o *ResizeOperation) String() string {
	// Validate operation
	if err := o.Validate(); err != nil {
		return ""
	}

	// Form return value
	str := OPERATION_NAME_RESIZE + QUERY_STRING_ENTRY_DELIMITER + "{w}" + DIMENSION_DELIMITER + "{h}"

	// Replace macros
	str = strings.Replace(str, "{w}", helpers.Int642String(o.dimensions.Width), -1)
	str = strings.Replace(str, "{h}", helpers.Int642String(o.dimensions.Height), -1)

	return str
}

// Validate returns a boolean indicating if the operation can be run,
// including checking source image against proposed operation parameters
func (o *ResizeOperation) Validate() error {
	// Verify dimensions exist
	if o.dimensions == nil {
		return fmt.Errorf("Invalid dimensions. Operation appears to not have been initialized")
	}

	// Verify target dimensions
	if o.dimensions.Width == 0 || o.dimensions.Height == 0 {
		return fmt.Errorf("Invalid target dimensions detected: %v", o.dimensions)
	}

	// Verify operation isn't trying to up-size image
	if o.dimensions.Width > o.mi.GetWidth() || o.dimensions.Height > o.mi.GetHeight() {
		return fmt.Errorf("Upsizing not supported for resize operations")
	}

	return nil
}

// parse is used to parse an operation's raw value and convert it
// into usable data for the operation
func (o *ResizeOperation) parse() error {
	var (
		// Error to be used throughout this method
		err error
	)

	// Split raw value and validation result
	b := strings.Split(o.rawValue, DIMENSION_DELIMITER)
	if len(b) != 2 {
		return fmt.Errorf("Invalid dimensions passed in for operation")
	}

	// Set operation dimensions
	o.dimensions, err = NewDimensionValues(b[0], b[1], o.mi.GetWidth(), o.mi.GetHeight())
	if err != nil {
		return err
	}

	return nil
}
