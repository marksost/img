// Tests the utils.go file
package operations

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("utils.go", func() {
	Describe("`NewDimensionValues` method", func() {
		var (
			// Input for `NewDimensionValues` input
			input []*NewDimensionValuesTestData
		)

		BeforeEach(func() {
			// Set input
			input = []*NewDimensionValuesTestData{
				// Width errors out
				&NewDimensionValuesTestData{
					Width:        "0.0xw",
					Height:       "*",
					ReturnsError: true,
				},
				// Height errors out
				&NewDimensionValuesTestData{
					Width:        "*",
					Height:       "0.0xh",
					ReturnsError: true,
				},
				// Two wildcards
				&NewDimensionValuesTestData{
					Width:        "*",
					Height:       "*",
					ReturnsError: true,
				},
				// Width ratio error
				&NewDimensionValuesTestData{
					Width:        "*",
					Height:       "100",
					SourceHeight: 0,
					ReturnsError: true,
				},
				// Height ratio error
				&NewDimensionValuesTestData{
					Width:        "100",
					Height:       "*",
					SourceWidth:  0,
					ReturnsError: true,
				},
				// Valid inputs
				// Both width/height
				&NewDimensionValuesTestData{
					Width:        "100", // pixels
					Height:       "100", // pixels
					SourceWidth:  200,
					SourceHeight: 300,
					ReturnWidth:  100,
					ReturnHeight: 100,
				},
				// Width wildcard
				&NewDimensionValuesTestData{
					Width:        "*",
					Height:       "100", // pixels
					SourceWidth:  200,
					SourceHeight: 300,
					ReturnWidth:  66,
					ReturnHeight: 100,
				},
				// Height wildcard
				&NewDimensionValuesTestData{
					Width:        "100", // pixels
					Height:       "*",
					SourceWidth:  200,
					SourceHeight: 300,
					ReturnWidth:  100,
					ReturnHeight: 150,
				},
				// Width relative value
				&NewDimensionValuesTestData{
					Width:        "0.5xw", // relative
					Height:       "100",   // pixels
					SourceWidth:  200,
					SourceHeight: 300,
					ReturnWidth:  100,
					ReturnHeight: 100,
				},
				// Height relative value
				&NewDimensionValuesTestData{
					Width:        "100",   // pixels
					Height:       "0.5xh", // relative
					SourceWidth:  200,
					SourceHeight: 300,
					ReturnWidth:  100,
					ReturnHeight: 150,
				},
			}
		})

		It("Returns either a set of valid dimensions or an error", func() {
			// Loop through test data
			for _, data := range input {
				// Call method
				dims, err := NewDimensionValues(data.Width, data.Height, data.SourceWidth, data.SourceHeight)

				// Verify return value
				if data.ReturnsError {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).To(Not(HaveOccurred()))
					Expect(dims.Width).To(Equal(data.ReturnWidth))
					Expect(dims.Height).To(Equal(data.ReturnHeight))
				}
			}
		})
	})

	Describe("`Dimension2Pixels` method", func() {
		var (
			// Input for `Dimension2Pixels` input
			input []*Dimension2PixelsTestData
		)

		BeforeEach(func() {
			// Set input
			input = []*Dimension2PixelsTestData{
				// Wildcard detected
				&Dimension2PixelsTestData{
					Dimension:      "*",
					ReturnWildcard: true,
				},
				// Ratio2Pixels error
				&Dimension2PixelsTestData{
					Dimension:    "0.0xw",
					ReturnsError: true,
				},
				// Ratio2Pixels without error
				&Dimension2PixelsTestData{
					Dimension:       "0.5xw",
					SourceDimension: 200,
					ReturnDimension: 100,
					ReturnWildcard:  false,
				},
				// String2Int64 without error
				&Dimension2PixelsTestData{
					Dimension:       "600",
					ReturnDimension: 600,
					ReturnWildcard:  false,
				},
				// Invalid data
				&Dimension2PixelsTestData{
					Dimension:    "foo",
					ReturnsError: true,
				},
			}
		})

		It("Returns a dimension, a widcard flag, and an error", func() {
			// Loop through test data
			for _, data := range input {
				// Call method
				dim, wc, err := Dimension2Pixels(data.Dimension, data.SourceDimension)

				// Verify return value
				if data.ReturnsError {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).To(Not(HaveOccurred()))
					Expect(dim).To(Equal(data.ReturnDimension))
					Expect(wc).To(Equal(data.ReturnWildcard))
				}
			}
		})
	})

	Describe("`Ratio2Pixels` method", func() {
		var (
			// Input for `Ratio2Pixels` input
			input []*Ratio2PixelsTestData
		)

		BeforeEach(func() {
			// Set input
			input = []*Ratio2PixelsTestData{
				// Zero-based ratio
				&Ratio2PixelsTestData{
					Dimension:    "0.0xw",
					ReturnsError: true,
				},
				// Valid inputs
				&Ratio2PixelsTestData{
					Dimension:       "0.5xw",
					SourceDimension: 600,
					ReturnDimension: 300,
				},
				&Ratio2PixelsTestData{
					Dimension:       "1.5xh",
					SourceDimension: 200,
					ReturnDimension: 300,
				},
				&Ratio2PixelsTestData{
					Dimension:       "0.25xh",
					SourceDimension: 400,
					ReturnDimension: 100,
				},
			}
		})

		It("Returns a dimension or an error", func() {
			// Loop through test data
			for _, data := range input {
				// Call method
				dim, err := Ratio2Pixels(data.Dimension, data.SourceDimension)

				// Verify return value
				if data.ReturnsError {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).To(Not(HaveOccurred()))
					Expect(dim).To(Equal(data.ReturnDimension))
				}
			}
		})
	})

	Describe("`RatioFromDimension` method", func() {
		var (
			// Input for `RatioFromDimension` input
			input []*RatioFromDimensionTestData
		)

		BeforeEach(func() {
			// Set input
			input = []*RatioFromDimensionTestData{
				// Zero-based denominator
				&RatioFromDimensionTestData{
					Denominator:  0,
					ReturnsError: true,
				},
				// Valid inputs
				&RatioFromDimensionTestData{
					Numerator:   10,
					Denominator: 5,
					ReturnRatio: 2,
				},
				&RatioFromDimensionTestData{
					Numerator:   25,
					Denominator: 5,
					ReturnRatio: 5,
				},
			}
		})

		It("Returns a dimension or an error", func() {
			// Loop through test data
			for _, data := range input {
				// Call method
				r, err := RatioFromDimension(data.Numerator, data.Denominator)

				// Verify return value
				if data.ReturnsError {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).To(Not(HaveOccurred()))
					Expect(r).To(Equal(data.ReturnRatio))
				}
			}
		})
	})
})
