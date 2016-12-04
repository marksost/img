package operations

import (
	// Standard lib
	"fmt"
	"strings"

	// Internal
	"github.com/marksost/img/image/mutableimages"
)

const (
	// The max number of operations allowed to be run per-request
	MAX_OPERATIONS = 5
	// The name of the crop operation
	OPERATION_NAME_CROP = "crop"
	// The name of the resize operation
	OPERATION_NAME_RESIZE = "resize"
	// The delimiter to be used when splitting query strings
	QUERY_STRING_DELIMITER = "&"
	// The delimiter to be used when splitting query string keys and values
	QUERY_STRING_ENTRY_DELIMITER = "="
)

type (
	// Inteface all image operations must satisfy
	Operation interface {
		// Public interface methods
		Process(*mutableimages.MutableImage) error
		Validate() error

		// Internal property methods
		String() string
	}
	// Struct representing an orchestrator for handling all image operations
	OperationController struct {
		// A slice of zero or more operations to run on an image
		Operations []Operation
		// A string representing the raw query string from the request
		queryString string
	}
)

// NewOperation creates a new operation and returns it
func NewOperation(operationType, value string) (Operation, error) {
	// Set default return value
	var op Operation

	switch operationType {
	case "crop":
		op = &CropOperation{rawValue: value}
	case "resize":
		op = &ResizeOperation{rawValue: value}
	default:
		return nil, fmt.Errorf("Unsupported operation type: %s", operationType)
	}

	return op, nil
}

// NewOperationController creates a new `OperationController` and returns it
func NewOperationController(qs []byte) *OperationController {
	// Create new operation controller
	oc := &OperationController{
		Operations: make([]Operation, 0, MAX_OPERATIONS),
		// NOTE: String conversion here may cause weirdness with non-UTF-8 chars
		queryString: string(qs),
	}

	// Filter URL params and set up operations if needed
	if oc.queryString != "" {
		oc.filterParams()
	}

	return oc
}

// Process takes a mutable image as input, iterates over each registered operation,
// and processes the image through the operation, returning an error if any occurs
func (oc *OperationController) Process(mi *mutableimages.MutableImage) error {
	// Loop through registered operations
	for _, op := range oc.Operations {
		if err := op.Process(mi); err != nil {
			return err
		}
	}

	return nil
}

// filterParams takes a raw query string from a request, splits it up
// into usable bits, validates each bit, and creates image operations
// from them when possible
func (oc *OperationController) filterParams() {
	// Reset operations slice
	oc.Operations = make([]Operation, 0, MAX_OPERATIONS)

	// Force lower-case for query string
	oc.queryString = strings.ToLower(oc.queryString)

	// Split query string on delimiter
	queries := strings.SplitN(oc.queryString, QUERY_STRING_DELIMITER, -1)

	// Loop through query string entries
	for _, query := range queries {
		// Verify max length hasn't been reached
		if len(oc.Operations) >= MAX_OPERATIONS {
			break
		}

		// Split query into key/value pairs
		bits := strings.Split(query, QUERY_STRING_ENTRY_DELIMITER)

		// Verify valid length and the operation is allowed
		if len(bits) != 2 {
			continue
		}

		// Create new operation
		operation, err := NewOperation(bits[0], bits[1])
		if err != nil {
			continue
		}

		// Append new operation to operations slice
		oc.Operations = append(oc.Operations, operation)
	}
}
