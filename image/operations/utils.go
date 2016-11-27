// operation util methods are used to help manipulate operation parameters
// like dimensions and crop values
package operations

import (
	// Standard lib
	"fmt"

	// Internal
	"github.com/marksost/img/helpers"
)

// NewDimensionValues takes width and height string values (gotten from a request)
// and converts them, using source dimensions if needed, into an ordered struct
// of width and height int64's for use within operations
func NewDimensionValues(w, h string, sw, sh int64) (*DimensionValues, error) {
	var (
		// Error to be used throughout this method
		err error
		// Width and height in pixels
		pWidth, pHeight int64
		// Ratio to be used when converting wildcard dimensions
		ratio float64
		// Booleans indicating if the width and height dimensions are "wildcards"
		wcWidth, wcHeight bool
	)

	// Get width in pixels
	if pWidth, wcWidth, err = Dimension2Pixels(w, sw); err != nil {
		return nil, err
	}

	// Get height in pixels
	if pHeight, wcHeight, err = Dimension2Pixels(h, sh); err != nil {
		return nil, err
	}

	// Check that max one dimension is a wildcard
	if wcWidth && wcHeight {
		return nil, fmt.Errorf("Only one dimension may be a wildcard")
	}

	// Handle width wildcards
	if wcWidth {
		// Get ratio
		if ratio, err = RatioFromDimension(pHeight, sh); err != nil {
			return nil, err
		}

		// Set width
		pWidth = int64(float64(sw) * ratio)
	}

	// Handle height wildcards
	if wcHeight {
		// Get ratio
		if ratio, err = RatioFromDimension(pWidth, sw); err != nil {
			return nil, err
		}

		// Set height
		pHeight = int64(float64(sh) * ratio)
	}

	return &DimensionValues{Width: pWidth, Height: pHeight}, nil
}

// Dimension2Pixels converts a single dimension (width or height)
// from a number of different formats into pixels when possible
// NOTE: The second return value indicates if the dimension is a "wildcard" or not
func Dimension2Pixels(dimension string, sourceDimension int64) (int64, bool, error) {
	// Check if dimension is a wildcard
	if dimension == DIMENSION_WILDCARD {
		return 0, true, nil
	}

	// Check if dimension is image-relative
	if len(dimension) > 2 && string(dimension[len(dimension)-2]) == "x" {
		// Get pixels from ratio
		px, err := Ratio2Pixels(dimension, sourceDimension)
		return px, false, err
	}

	// Convert string to int64
	// NOTE: String2Int64 abstracts conversion errors away, thus the extra value check
	if px := helpers.String2Int64(dimension); px != 0 {
		return px, false, nil
	}

	// Return fall-through error
	return 0, false, fmt.Errorf("Invalid dimension detected: %s", dimension)
}

// Ratio2Pixels converts a "ratio" dimension (ex: 1.234xw)
// into pixles when possible
func Ratio2Pixels(dimension string, sourceDimension int64) (int64, error) {
	// Parse ratio as a float
	r := helpers.String2Float64(dimension[0 : len(dimension)-2])

	// Check for valid input
	if r == 0.0 {
		return 0, fmt.Errorf("A ratio must be a non-zero value: %s", dimension)
	}

	return int64(r * float64(sourceDimension)), nil
}

// RatioFromDimension returns a ratio of one integer to another
func RatioFromDimension(numerator, denominator int64) (float64, error) {
	// check for valid input
	if denominator == 0 {
		return 0.0, fmt.Errorf("Denominator must be a non-zero value: %q", denominator)
	}

	return float64(numerator) / float64(denominator), nil
}
