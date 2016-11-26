// Tests the image.go file
package image

import (
	// Standard lib
	"io/ioutil"
	"path"

	// Internal
	"github.com/marksost/img/image/mutableimages"
	"github.com/marksost/img/image/utils"

	// Third-party
	"github.com/kataras/iris"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

var _ = Describe("image.go", func() {
	var (
		// Mock iris context to use within tests
		ctx *iris.Context
		// Mock image to test
		i *Image
	)

	BeforeEach(func() {
		// Create mock context
		ctx = &iris.Context{
			RequestCtx: &fasthttp.RequestCtx{},
		}

		// Create mock image
		i = NewImage(ctx)
	})

	Describe("`NewImage` method", func() {
		It("Returns a valid image", func() {
			// Call method
			i := NewImage(ctx)

			// Verify image was properly created and returned
			Expect(i).To(Not(BeNil()))
		})
	})

	Describe("Image internal property methods", func() {
		Describe("`Data` method", func() {
			BeforeEach(func() {
				// Reset data
				data, err := ioutil.ReadFile(path.Join("../test/images/1x1.gif"))
				if err != nil {
					panic("Error reading image. Tests cannot continue. " + err.Error())
				}

				// Create mi
				mi, err := mutableimages.NewMutableImage(data, utils.GIF_MIME)
				if err != nil {
					panic("Error creating mutable image. Tests cannot continue. " + err.Error())
				}

				// Set new utility structs to ensure predictable values
				i.utils = &ImageUtils{
					MutableImage: mi,
				}
			})

			It("Returns a byte slice of data", func() {
				// Call method
				data := i.Data()

				// Verify return value
				Expect(len(data)).To(Not(Equal(0)))
			})
		})
	})

	Describe("Image utils proxy methods", func() {
		BeforeEach(func() {
			// Set new utility structs to ensure predictable values
			i.utils = &ImageUtils{
				Downloader: utils.NewDownloader("/foo-url.com/path/to/image.jpg"),
			}
		})

		Describe("`MimeType` method", func() {
			It("Returns a default MIME type", func() {
				// Call method
				mimeType := i.MimeType()

				// Verify return value
				Expect(mimeType).To(Equal(utils.DEFAULT_MIME_TYPE))
			})
		})

		Describe("`RawData` method", func() {
			It("Returns an empty byte slice", func() {
				// Call method
				data := i.RawData()

				// Verify return value
				Expect(len(data)).To(Equal(0))
			})
		})

		Describe("`Url` method", func() {
			It("Returns a `url.URL` instance", func() {
				// Call method
				u := i.Url()

				// Verify return value
				Expect(u.String()).To(Equal("http://foo-url.com/path/to/image.jpg"))
			})
		})
	})
})
