// error defines all request error functionality to be used to handle errors
// that occur while processing images within this package
package image

type (
	// Struct representing an error that occurred during a image processing request
	// NOTE: `ImageRequestError` satisfies the standard `error` interface,
	// and can be used interchangably
	// A cast (like: `err.(*image.ImageRequestError)`) is needed
	// to access non-nterface methods
	ImageRequestError struct {
		code int    // The error code to return for the error
		str  string // The error string
	}
)

// NewError creates a new `ImageRequestError` and returns it
func NewError(code int, str string) *ImageRequestError {
	// Create and return new error
	return &ImageRequestError{code: code, str: str}
}

// Code returns the internal `code` property of the error
func (e *ImageRequestError) Code() int {
	return e.code
}

// Error returns the internal `str` property of the error
func (e *ImageRequestError) Error() string {
	return e.str
}
