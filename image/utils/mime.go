// mime contains all information and functionality around determining the MIME type
// of a requested image
package utils

import (
	// Standard lib
	"bytes"
)

const (
	// The default MIME type to return when no MIME type was set from the downloaded image
	DEFAULT_MIME_TYPE = "application/octet-stream"

	// MIME types for all supported image formats
	GIF_MIME  = "image/gif"
	JPEG_MIME = "image/jpeg"
	PNG_MIME  = "image/png"
	TIFF_MIME = "image/tiff"
)

var (
	// Map of MIME types to the first two bytes of their data
	// NOTE: Used to identify what type of image is being requested
	// NOTE: Different image types all have consistent starting byte patterns
	MimeTypes = map[string][]byte{
		GIF_MIME:  []byte{0x47, 0x49},
		JPEG_MIME: []byte{0xff, 0xd8},
		PNG_MIME:  []byte{0x89, 0x50},
		TIFF_MIME: []byte{0x49, 0x49},
	}
)

// getMimeType attempts to determine the correct MIME type for a given byte slice
// Will return a default MIME type when no match is found
func getMimeType(data []byte) string {
	// Loop through defined types, checking the first two bytes of the input
	for mime, slice := range MimeTypes {
		if bytes.Equal(data[:2], slice) {
			return mime
		}
	}

	return DEFAULT_MIME_TYPE
}
