// Tests the gif.go file
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

var _ = Describe("gif.go", func() {
	var (
		// Byte slice of image data to use throughout testing
		data []byte
		// Error to use throughout testing
		err error
		// Mock GIF mutable image to test
		mi *GifMutableImage
		// Mock processable image to use throughout testing
		pi *ProcessableImage
	)

	BeforeEach(func() {
		// Initalize config instance
		config.Init()

		// Set data
		data, err = ioutil.ReadFile(path.Join("../../test/images/1x1.gif"))
		if err != nil {
			panic("Error reading image. Tests cannot continue. " + err.Error())
		}

		// Create processable image
		pi = &ProcessableImage{
			Data:      data,
			ImageType: utils.GIF_MIME,
		}

		// Create static mutable image
		mi, err = NewGifMutableImage(pi)
		if err != nil {
			panic("Error creating static mutable image. Tests cannot continue. " + err.Error())
		}
	})

	Describe("`NewGifMutableImage` method", func() {
		It("Returns a valid GifMutableImage instance", func() {
			// Call method
			mi, err := NewGifMutableImage(pi)

			// Verify return value
			Expect(err).To(Not(HaveOccurred()))
			Expect(mi.GetWidth()).To(BeEquivalentTo(1))  // NOTE: Equiv because of int vs int64
			Expect(mi.GetHeight()).To(BeEquivalentTo(1)) // NOTE: Equiv because of int vs int64
		})
	})

	Describe("MutableImage interface methods", func() {
		Describe("`GetWidth` method", func() {
			It("Returns the width of the image", func() {
				// Call method
				w := mi.GetWidth()

				// Verify return value
				Expect(w).To(Equal(int64(1)))
			})
		})

		Describe("`GetHeight` method", func() {
			It("Returns the height of the image", func() {
				// Call method
				h := mi.GetHeight()

				// Verify return value
				Expect(h).To(Equal(int64(1)))
			})
		})

		Describe("`Img` method", func() {
			It("Returns an internal `img` property", func() {
				// Call method
				img := mi.Img()

				// Verify return value
				Expect(len(img.Data)).To(Equal(len(data)))
				Expect(img.ImageType).To(Equal(utils.GIF_MIME))
			})
		})

		Describe("`SetDefaults` method", func() {
			It("Sets default values for internal properties", func() {
				// NOTE: No defaults set as of now
				// Calling method to ensure it's reachable
				mi.SetDefaults()
			})
		})

		Describe("`SetDimensions` method", func() {
			BeforeEach(func() {
				// Reset width and height
				mi.width = 0
				mi.height = 0
			})

			It("Sets image dimensions", func() {
				// Call method
				mi.SetDimensions()

				// Verify dimensions were not set
				Expect(mi.width).To(BeEquivalentTo(1))  // NOTE: Equiv because of int vs int64
				Expect(mi.height).To(BeEquivalentTo(1)) // NOTE: Equiv because of int vs int64
			})
		})
	})
})
