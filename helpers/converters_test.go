// Tests the converters.go file
package helpers

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("converters.go", func() {
	Describe("`Bool2String` method", func() {
		var (
			// Input for `Bool2String` input
			input map[bool]string
		)

		BeforeEach(func() {
			// Set input
			input = map[bool]string{
				true:  "true",
				false: "false",
			}
		})

		It("Converts a bool to a string", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := Bool2String(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})

	Describe("`Float642String` method", func() {
		var (
			// Input for `Float642String` input
			input map[float64]string
		)

		BeforeEach(func() {
			// Set input
			input = map[float64]string{
				0:       "0",
				1.0:     "1",
				234.567: "234.567",
			}
		})

		It("Converts a float64 to a string", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := Float642String(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})

	Describe("`Int2String` method", func() {
		var (
			// Input for `Int2String` input
			input map[int]string
		)

		BeforeEach(func() {
			// Set input
			input = map[int]string{
				0:   "0",
				1:   "1",
				234: "234",
			}
		})

		It("Converts an int to a string", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := Int2String(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})

	Describe("`Int642String` method", func() {
		var (
			// Input for `Int642String` input
			input map[int64]string
		)

		BeforeEach(func() {
			// Set input
			input = map[int64]string{
				0:   "0",
				1:   "1",
				234: "234",
			}
		})

		It("Converts an int64 to a string", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := Int642String(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})

	Describe("`Interface2String` method", func() {
		var (
			// Input for `Interface2String` input
			input map[interface{}]string
		)

		BeforeEach(func() {
			// Set input
			input = map[interface{}]string{
				234.567:     "234.567",
				0:           "0",
				int64(1234): "1234",
				"foo":       "foo",
				true:        "",
			}
		})

		It("Converts an interface to a string", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := Interface2String(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})

	Describe("`MapFromInterface` method", func() {
		var (
			// Input for `MapFromInterface` input
			input map[string]interface{}
		)

		BeforeEach(func() {
			// Set input
			input = map[string]interface{}{"foo": 1234}
		})

		It("Converts an interface to a `map[string]interface{}`", func() {
			// Call method
			actual := MapFromInterface(input)

			// Verify return value
			// NOTE: Verifies map values are interfaces
			Expect(actual["foo"].(int)).To(Equal(1234))
		})
	})

	Describe("`String2Int` method", func() {
		var (
			// Input for `String2Int` input
			input map[string]int
		)

		BeforeEach(func() {
			// Set input
			input = map[string]int{
				"foo":  0,
				"1234": 1234,
			}
		})

		It("Converts a string to an int", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := String2Int(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})

	Describe("`String2Int64` method", func() {
		var (
			// Input for `String2Int64` input
			input map[string]int64
		)

		BeforeEach(func() {
			// Set input
			input = map[string]int64{
				"foo":  0,
				"1234": 1234,
			}
		})

		It("Converts a string to an int64", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := String2Int64(input)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})
})
