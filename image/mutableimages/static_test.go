// Tests the static.go file
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

var _ = Describe("static.go", func() {
	var (
		// Byte slice of image data to use throughout testing
		data []byte
		// Error to use throughout testing
		err error
		// Mock static mutable image to test
		mi *StaticMutableImage
		// Mock processable image to use throughout testing
		pi *ProcessableImage
	)

	BeforeEach(func() {
		// Initalize config instance
		config.Init()

		// Set data
		data, err = ioutil.ReadFile(path.Join("../../test/images/1x1.jpg"))
		if err != nil {
			panic("Error reading image. Tests cannot continue. " + err.Error())
		}

		// Create processable image
		pi = &ProcessableImage{
			Data:      data,
			ImageType: utils.JPEG_MIME,
		}

		// Create static mutable image
		mi, err = NewStaticMutableImage(pi)
		if err != nil {
			panic("Error creating static mutable image. Tests cannot continue. " + err.Error())
		}
	})

	Describe("`NewStaticMutableImage` method", func() {
		It("Returns a valid StaticMutableImage instance", func() {
			// Call method
			mi, err := NewStaticMutableImage(pi)

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
				Expect(img.ImageType).To(Equal(utils.JPEG_MIME))
			})
		})

		Describe("`SetDefaults` method", func() {
			It("Sets default values for internal properties", func() {
				// Call method
				mi.SetDefaults()

				// Verify defaults were set
				Expect(mi.quality).To(Equal(config.GetInstance().Images.DefaultQuality))
			})
		})

		Describe("`SetDimensions` method", func() {
			Context("With invalid image data", func() {
				BeforeEach(func() {
					// Reset image data
					pi.Data = make([]byte, 0)

					// Reset width and height
					mi.width = 0
					mi.height = 0
				})

				It("Returns early", func() {
					// Call method
					mi.SetDimensions()

					// Verify dimensions were not set
					Expect(mi.width).To(BeEquivalentTo(0))  // NOTE: Equiv because of int vs int64
					Expect(mi.height).To(BeEquivalentTo(0)) // NOTE: Equiv because of int vs int64
				})
			})

			Context("With valid image data", func() {
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
})
