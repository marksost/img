// Tests the operations.go file
package operations

import (
	// Standard lib
	"io/ioutil"
	"path"
	"strings"

	// Internal
	"github.com/marksost/img/config"
	"github.com/marksost/img/image/mutableimages"
	"github.com/marksost/img/image/utils"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("operations.go", func() {
	var (
		// Byte slice of image data to use throughout testing
		data []byte
		// Error to use throughout testing
		err error
		// Byte slice simulating a request query string to use throughout testing
		qs []byte
		// Mock static mutable image to use throughout testing
		mi mutableimages.MutableImage
		// Mock operation controller to test
		oc *OperationController
		// String to use for quest strings
		str string
	)

	BeforeEach(func() {
		// Initalize config instance
		config.Init()

		// Set query string
		str = "foo=bar&test=one&single-key"
		qs = []byte(str)

		// Create mock operation controller
		oc = NewOperationController(qs)

		// Set data
		data, err = ioutil.ReadFile(path.Join("../../test/images/1x1.jpg"))
		if err != nil {
			panic("Error reading image. Tests cannot continue. " + err.Error())
		}

		// Create static mutable image
		mi, err = mutableimages.NewMutableImage(data, utils.JPEG_MIME)
		if err != nil {
			panic("Error creating static mutable image. Tests cannot continue. " + err.Error())
		}
	})

	Describe("`NewOperation` method", func() {
		Context("With an invalid operation type", func() {
			It("Returns an error", func() {
				// Call method
				_, err := NewOperation("foo", "bar")

				// Verify return value
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With a valid operation type", func() {
			It("Returns an operation", func() {
				// Call method
				op, err := NewOperation("resize", "bar")

				// Verify return value
				Expect(err).To(Not(HaveOccurred()))
				Expect(op.String()).To(BeAssignableToTypeOf("string"))
			})
		})
	})

	Describe("`NewOperationController` method", func() {
		It("Returns a valid OperationController instance", func() {
			// Call method
			oc := NewOperationController(qs)

			// Verify return value
			Expect(oc).To(Not(BeNil()))
			Expect(oc.queryString).To(Equal(string(qs)))
		})
	})

	Describe("OperationController public functionality methods", func() {
		Describe("`Process` method", func() {
			Context("With an operation that returns an error", func() {
				BeforeEach(func() {
					// Set operations
					oc.Operations = []Operation{
						&MockOperationWithError{},
					}
				})

				It("Returns the operation's erorr", func() {
					// Call method
					err := oc.Process(&mi)

					// Verify return value
					Expect(err).To(HaveOccurred())
				})
			})

			Context("With an operation that returns an error", func() {
				BeforeEach(func() {
					// Set operations
					oc.Operations = []Operation{
						&MockOperationWithoutError{},
					}
				})

				It("Returns nil, after processing all operations", func() {
					// Call method
					err := oc.Process(&mi)

					// Verify return value
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})
	})

	Describe("OperationController utility methods", func() {
		Describe("`filterParams` method", func() {
			BeforeEach(func() {
				// Set query string
				// NOTE: Tests various cases: max operations, invalid params, unrecognized operations
				str = "resize=foo&RESIZE=bar&resize=baz&test=one&single-key&resize=two&resize=three&resize=four"
				oc.queryString = str
			})

			It("Formats a query string and sets up operations", func() {
				// Call method
				oc.filterParams()

				// Verify query string was set to lower case
				Expect(oc.queryString).To(Equal(strings.ToLower(str)))

				// Verify length of operations
				Expect(len(oc.Operations)).To(Equal(MAX_OPERATIONS))
			})
		})
	})
})
