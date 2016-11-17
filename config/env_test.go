// Tests the env.go file
package config

import (
	// Standard lib
	"flag"
	"os"
	"reflect"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("env.go", func() {
	var (
		// Mock application config to test
		c *Config
	)

	BeforeEach(func() {
		// Set test environment variables
		os.Setenv(ENV_PREFIX+"HANDLE_FIELDS_FOO", "foo")
		os.Setenv(ENV_PREFIX+"HANDLE_FIELDS_BAR", "1234")
		os.Setenv(ENV_PREFIX+"HANDLE_FIELDS_BAZ", "1")
		os.Setenv(ENV_PREFIX+"HANDLE_FIELDS_TEST_FOO", "test-foo")

		// Create config instance
		c = &Config{}
	})

	Describe("Config struct methods", func() {
		Describe("`handleFields` method", func() {
			var (
				// Input for `handleFields` input
				input *HandleFieldsTestData
			)

			BeforeEach(func() {
				// Set input
				input = &HandleFieldsTestData{}
			})

			It("Sets configuration values based on environment variable values and types", func() {
				// Call method
				c.handleFields(reflect.ValueOf(input))

				// Verify values were set
				Expect(input.Foo).To(Equal("foo"))
				Expect(input.Bar).To(Equal(1234))
				Expect(input.Baz).To(BeTrue())
				Expect(input.Test.Foo).To(Equal("test-foo"))
			})

			It("Sets flags based on environment variables that set set", func() {
				// Call method
				c.handleFields(reflect.ValueOf(input))

				// Verify flags were set
				Expect(flag.Lookup("handle-fields-foo")).To(Not(BeNil()))
				Expect(flag.Lookup("handle-fields-bar")).To(Not(BeNil()))
				Expect(flag.Lookup("handle-fields-baz")).To(Not(BeNil()))
				Expect(flag.Lookup("handle-fields-test-foo")).To(Not(BeNil()))
			})
		})

		Describe("`formFlagName` method", func() {
			var (
				// Input for `formFlagName` input
				input map[string]string
			)

			BeforeEach(func() {
				// Set input
				input = map[string]string{
					"test":        "test",
					"TEST":        "test",
					"test_foo":    "test-foo",
					"img_foo_bar": "foo-bar",
					"IMG_FOO_BAR": "foo-bar",
				}
			})

			It("Returns a formatted flag name", func() {
				// Loop through test data
				for input, expected := range input {
					// Call method
					actual := c.formFlagName(input)

					// Verify result
					Expect(actual).To(Equal(expected))
				}
			})
		})
	})
})
