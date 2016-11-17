// Tests the helpers.go file
package helpers

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("helpers.go", func() {
	Describe("`SliceContains` method", func() {
		var (
			// Input for `SliceContains` input
			input map[*SliceContainsTestData]bool
		)

		BeforeEach(func() {
			// Set input
			input = map[*SliceContainsTestData]bool{
				&SliceContainsTestData{"foo", []string{}}:                    false,
				&SliceContainsTestData{"foo", []string{"bar", "baz"}}:        false,
				&SliceContainsTestData{"foo", []string{"foo", "bar", "baz"}}: true,
			}
		})

		It("Returns a boolean indicating if a slice of strings contains a specific string", func() {
			// Loop through test data
			for input, expected := range input {
				// Call method
				actual := SliceContains(input.Needle, input.Haystack)

				// Verify return value
				Expect(actual).To(Equal(expected))
			}
		})
	})
})
