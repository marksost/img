// Tests the image.go file
package image

import (
	// Third-party
	"github.com/kataras/iris"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("image.go", func() {
	var (
		// Mock image to test
		i *Image
	)

	BeforeEach(func() {
		// Create mock image
		i = NewImage(&iris.Context{})
	})

	Describe("`NewImage` method", func() {
		It("Returns a valid image", func() {
			// Call method
			i := NewImage(&iris.Context{})

			// Verify image was properly created and returned
			Expect(i).To(Not(BeNil()))
			Expect(i.MimeType()).To(Equal(DEFAULT_MIME_TYPE))
		})
	})

	Describe("Image internal property methods", func() {
		Describe("`Data` method", func() {
			BeforeEach(func() {
				// Set data
				i.outputData = []byte("this is some test data")
			})

			It("Returns a byte slice of data", func() {
				// Call method
				data := i.Data()

				// Verify return value
				Expect(string(data)).To(Equal("this is some test data"))
			})
		})

		Describe("`MimeType` method", func() {
			Context("With no MIME type set for the image", func() {
				BeforeEach(func() {
					// Set empty MIME type
					i.mimeType = ""
				})

				It("Returns a default MIME type", func() {
					// Call method
					mimeType := i.MimeType()

					// Verify return value
					Expect(mimeType).To(Equal(DEFAULT_MIME_TYPE))
				})
			})

			Context("With a MIME type set for the image", func() {
				BeforeEach(func() {
					// Set non-empty MIME type
					i.mimeType = "foo-mime"
				})

				It("Returns a the set MIME type", func() {
					// Call method
					mimeType := i.MimeType()

					// Verify return value
					Expect(mimeType).To(Equal("foo-mime"))
				})
			})
		})
	})
})
