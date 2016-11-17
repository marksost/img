// Tests the config.go file
package config

import (
	// Standard lib
	"os"
	"path"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("config.go", func() {
	var (
		// Mock application config to test
		c *Config
	)

	BeforeEach(func() {
		// Ensure config env var is unset
		os.Unsetenv(CONFIG_LOCATION)

		// Create config instance
		c = &Config{}
	})

	Describe("Config struct methods", func() {
		Describe("`readConfigFile` method", func() {
			Context("When no config location environment variable is set", func() {
				BeforeEach(func() {
					// Ensure config env var is unset
					os.Unsetenv(CONFIG_LOCATION)
				})

				It("Returns false", func() {
					// Verify return value
					Expect(c.readConfigFile()).To(BeFalse())
				})
			})

			Context("When the configuration file doesn't exist", func() {
				BeforeEach(func() {
					// Set config env var
					os.Setenv(CONFIG_LOCATION, path.Join("../test/data/doesnt-exist.json"))
				})

				It("Returns false", func() {
					// Verify return value
					Expect(c.readConfigFile()).To(BeFalse())
				})
			})

			Context("When the configuration file contains invalid JSON", func() {
				BeforeEach(func() {
					// Set config env var
					os.Setenv(CONFIG_LOCATION, path.Join("../test/data/invalid-config.json"))
				})

				It("Returns false", func() {
					// Verify return value
					Expect(c.readConfigFile()).To(BeFalse())
				})
			})

			Context("When the configuration file contains valid JSON", func() {
				BeforeEach(func() {
					// Set config env var
					os.Setenv(CONFIG_LOCATION, path.Join("../test/data/valid-config.json"))
				})

				It("Reads in the configuration, sets values, and returns true", func() {
					// Verify return value
					Expect(c.readConfigFile()).To(BeTrue())

					// Verify values were set
					Expect(c.Name).To(Equal("foo"))
				})
			})
		})
	})
})
