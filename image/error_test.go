// Tests the error.go file
package image

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("error.go", func() {
	var (
		// Mock error to test
		err *ImageRequestError
	)

	BeforeEach(func() {
		// Create mock error
		err = NewError(200, "This is a test error")
	})

	Describe("`NewError` method", func() {
		It("Returns a valid error", func() {
			// Call method
			err := NewError(1234, "mock-error-string")

			// Verify error was properly created and returned
			Expect(err).To(Not(BeNil()))
			Expect(err.Code()).To(Equal(1234))
			Expect(err.Error()).To(Equal("mock-error-string"))
		})
	})

	Describe("ImageRequestError struct methods", func() {
		Describe("`Code` method", func() {
			It("Returns the internal `code` property of the error", func() {
				// Call method
				code := err.Code()

				// Verify return value
				Expect(code).To(Equal(200))
			})
		})

		Describe("`Error` method", func() {
			It("Returns the internal `str` property of the error", func() {
				// Call method
				str := err.Error()

				// Verify return value
				Expect(str).To(Equal("This is a test error"))
			})
		})
	})
})
