// Tests the downloader.go file
package utils

import (
	// Standard lib
	"net/url"

	// Internal
	"github.com/marksost/img/helpers"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("downloader.go", func() {
	var (
		// Mock downloader to test
		d *Downloader
	)

	BeforeEach(func() {
		// Create mock downloader
		d = NewDownloader("/foo-url.com/path/to/image.jpg")
	})

	Describe("`NewDownloader` method", func() {
		It("Sets an populated URL instance and returns a valid downloader instance", func() {
			// Call method
			d := NewDownloader("/foo-url.com/path/to/image.jpg")

			// Verify downloader was properly created and returned
			Expect(d).To(Not(BeNil()))
			Expect(d.url.String()).To(Equal("http://foo-url.com/path/to/image.jpg"))
		})
	})

	Describe("Downloader public functionality methods", func() {
		Describe("`Download` method", func() {
			Context("When an `http.Get` error occurs", func() {
				BeforeEach(func() {
					// Set url
					d.url = &url.URL{Opaque: ":"}
				})

				It("Returns an error", func() {
					// Call method
					err := d.Download()

					// Verify return value
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When the response status code is not 200", func() {
				BeforeEach(func() {
					// Set url
					d.url, _ = url.Parse(helpers.GetMockServer("bad-request").URL)
				})

				It("Returns an error", func() {
					// Call method
					err := d.Download()

					// Verify return value
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When the response is successful", func() {
				BeforeEach(func() {
					// Set url
					d.url, _ = url.Parse(helpers.GetMockServer("default").URL)
				})

				It("Set the response body as data and returns no error", func() {
					// Call method
					err := d.Download()

					// Verify return value
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})
	})

	Describe("Downloader internal property methods", func() {
		Describe("`Data` method", func() {
			BeforeEach(func() {
				// Set data
				d.data = []byte("this is some test data")
			})

			It("Returns a byte slice of data", func() {
				// Call method
				data := d.Data()

				// Verify return value
				Expect(string(data)).To(Equal("this is some test data"))
			})
		})

		Describe("`MimeType` method", func() {
			Context("With no MIME type set", func() {
				BeforeEach(func() {
					// Set empty MIME type
					d.mimeType = ""
				})

				It("Returns a default MIME type", func() {
					// Call method
					mimeType := d.MimeType()

					// Verify return value
					Expect(mimeType).To(Equal(DEFAULT_MIME_TYPE))
				})
			})

			Context("With a MIME type set", func() {
				BeforeEach(func() {
					// Set non-empty MIME type
					d.mimeType = "foo-mime"
				})

				It("Returns a the set MIME type", func() {
					// Call method
					mimeType := d.MimeType()

					// Verify return value
					Expect(mimeType).To(Equal("foo-mime"))
				})
			})
		})

		Describe("`Url` method", func() {
			BeforeEach(func() {
				// Set url
				d.url, _ = url.Parse("http://foo-url.com/test")
			})

			It("Returns a `url.URL` instance", func() {
				// Call method
				u := d.Url()

				// Verify return value
				Expect(u.String()).To(Equal("http://foo-url.com/test"))
			})
		})
	})
})
