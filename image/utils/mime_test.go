// Tests the mime.go file
package utils

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mime.go", func() {
	Describe("`getMimeType` method", func() {
		var (
			// Input for `getMimeType` input
			input map[string][]byte
		)

		BeforeEach(func() {
			// Set input
			input = map[string][]byte{
				"application/octet-stream": []byte{0x99, 0x99}, // NOTE: Tests fallback
				"image/gif":                []byte{0x47, 0x49},
				"image/jpeg":               []byte{0xff, 0xd8},
				"image/png":                []byte{0x89, 0x50},
				"image/tiff":               []byte{0x49, 0x49},
			}
		})

		It("Returns either a matching MIME type or a default", func() {
			// Loop through test data
			for expected, input := range input {
				// Call method
				actual := getMimeType(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})
})
