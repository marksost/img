package operations

import (
	// Standard lib
	"fmt"
	"strings"

	// Internal
	"github.com/marksost/img/config"
	"github.com/marksost/img/helpers"
	"github.com/marksost/img/image/mutableimages"
)

type (
	// Struct representing a quality operation to be performed on an image
	QualityOperation struct {
		// Mutable image to use when processing this operation
		mi mutableimages.MutableImage
		// Raw query string value for this operation
		rawValue string
		// Value used when operating on the image
		value int64
	}
)

// Process is used to perform the actual operation processing
// on a given image
func (o *QualityOperation) Process(mi *mutableimages.MutableImage) error {
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

	// Return value from quality operation
	return o.mi.Quality(o.value)
}

// String returns a string representation of this operation
func (o *QualityOperation) String() string {
	// Validate operation
	if err := o.Validate(); err != nil {
		return ""
	}

	// Form return value
	str := OPERATION_NAME_QUALITY + QUERY_STRING_ENTRY_DELIMITER + "{q}"

	// Replace macros
	str = strings.Replace(str, "{q}", helpers.Int642String(o.value), -1)

	return str
}

// Validate returns a boolean indicating if the operation can be run,
// including checking source image against proposed operation parameters
func (o *QualityOperation) Validate() error {
	// Verify value is within range
	if o.value <= 0 || o.value > 100 {
		return fmt.Errorf("Invalid value")
	}

	return nil
}

// parse is used to parse an operation's raw value and convert it
// into usable data for the operation
func (o *QualityOperation) parse() error {
	// Convert raw value to int64
	o.value = helpers.String2Int64(o.rawValue)

	// Reset value if needed
	if err := o.Validate(); err != nil {
		o.value = int64(config.GetInstance().Images.DefaultQuality)
	}

	return nil
}
