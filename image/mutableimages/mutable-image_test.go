// Tests the mutable-image.go file
package mutableimages

import (
	// Standard lib
	"io/ioutil"
	"path"

	// Internal
	"github.com/marksost/img/config"
	"github.com/marksost/img/image/utils"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mutable-image.go", func() {
	var (
		// Byte slice of image data to use throughout testing
		data []byte
		// Error to use throughout testing
		err error
	)

	BeforeEach(func() {
		// Initalize config instance
		config.Init()

		// Set empty data
		data = make([]byte, 0)
	})

	Describe("`NewMutableImage` method", func() {
		Context("With an invalid type as input", func() {
			It("Returns an error", func() {
				// Call method
				_, err := NewMutableImage(data, "foo")

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With an invalid GIF as input data", func() {
			It("Returns an error", func() {
				// Call method
				_, err := NewMutableImage(data, utils.GIF_MIME)

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With a valid GIF as input data", func() {
			BeforeEach(func() {
				// Reset data
				data, err = ioutil.ReadFile(path.Join("../../test/images/1x1.gif"))
				if err != nil {
					panic("Error reading image. Tests cannot continue. " + err.Error())
				}
			})

			It("Returns a new static mutable image struct", func() {
				// Call method
				mi, err := NewMutableImage(data, utils.GIF_MIME)

				// Verify return value
				Expect(err).To(Not(HaveOccurred()))
				Expect(mi.GetWidth()).To(BeEquivalentTo(1))  // NOTE: Equiv because of int vs int64
				Expect(mi.GetHeight()).To(BeEquivalentTo(1)) // NOTE: Equiv because of int vs int64
			})
		})

		Context("With a JPEG as input data", func() {
			BeforeEach(func() {
				// Reset data
				data, err = ioutil.ReadFile(path.Join("../../test/images/1x1.jpg"))
				if err != nil {
					panic("Error reading image. Tests cannot continue. " + err.Error())
				}
			})

			It("Returns a new static mutable image struct", func() {
				// Call method
				mi, err := NewMutableImage(data, utils.JPEG_MIME)

				// Verify return value
				Expect(err).To(Not(HaveOccurred()))
				Expect(mi.GetWidth()).To(BeEquivalentTo(1))  // NOTE: Equiv because of int vs int64
				Expect(mi.GetHeight()).To(BeEquivalentTo(1)) // NOTE: Equiv because of int vs int64
			})
		})
	})
})
