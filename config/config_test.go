// Tests the config.go file
package config

import (
	// Standard lib
	"os"

	// Third-party
	log "github.com/Sirupsen/logrus"
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

		// Initialize config instance
		// NOTE: Provided for method test coverage
		c.Init()
	})

	Describe("Config struct methods", func() {
		Describe("Setter methods", func() {
			Describe("`setDefaults` method", func() {
				It("Sets default configuration values", func() {
					// Verify default values were set
					Expect(c.Server.Timeouts.Read).To(Equal(30))
				})
			})

			Describe("`setLoggerSettings` method", func() {
				var (
					// Input for `setLoggerSettings` input
					input map[string]log.Level
				)

				BeforeEach(func() {
					// Set logger formatter
					// NOTE: Provides switch statement coverage
					c.Log.Formatter = "json"

					// Set input
					input = map[string]log.Level{
						"error":   log.ErrorLevel,
						"fatal":   log.FatalLevel,
						"info":    log.InfoLevel,
						"invalid": log.DebugLevel,
						"panic":   log.PanicLevel,
						"warn":    log.WarnLevel,
					}
				})

				It("Sets logger settings", func() {
					// Loop through test data
					for level, expected := range input {
						// Set logger level
						c.Log.Level = level

						// Call method
						c.setLoggerSettings()

						// Verify level was set
						Expect(log.GetLevel()).To(Equal(expected))
					}
				})
			})
		})

		Describe("Environment checker methods", func() {
			Describe("`IsDevelopment` method", func() {
				Context("The environment is 'dev'", func() {
					BeforeEach(func() {
						// Set current environment
						c.Environment = ENV_DEVELOPMENT
					})

					It("Returns true", func() {
						// Verify return value
						Expect(c.IsDevelopment()).To(BeTrue())
					})
				})

				Context("The environment is not 'dev'", func() {
					BeforeEach(func() {
						// Set current environment
						c.Environment = ENV_TESTING
					})

					It("Returns false", func() {
						// Verify return value
						Expect(c.IsDevelopment()).To(BeFalse())
					})
				})
			})

			Describe("`IsProduction` method", func() {
				Context("The environment is 'prod'", func() {
					BeforeEach(func() {
						// Set current environment
						c.Environment = ENV_PRODUCTION
					})

					It("Returns true", func() {
						// Verify return value
						Expect(c.IsProduction()).To(BeTrue())
					})
				})

				Context("The environment is not 'prod'", func() {
					BeforeEach(func() {
						// Set current environment
						c.Environment = ENV_DEVELOPMENT
					})

					It("Returns false", func() {
						// Verify return value
						Expect(c.IsProduction()).To(BeFalse())
					})
				})
			})

			Describe("`IsTesting` method", func() {
				Context("The environment is 'test'", func() {
					BeforeEach(func() {
						// Set current environment
						c.Environment = ENV_TESTING
					})

					It("Returns true", func() {
						// Verify return value
						Expect(c.IsTesting()).To(BeTrue())
					})
				})

				Context("The environment is not 'test'", func() {
					BeforeEach(func() {
						// Set current environment
						c.Environment = ENV_DEVELOPMENT
					})

					It("Returns false", func() {
						// Verify return value
						Expect(c.IsTesting()).To(BeFalse())
					})
				})
			})
		})
	})

	Describe("`GetInstance` method", func() {
		It("Returns an instance of the initialized configuration struct", func() {
			// Call `Init` method to set up config
			Init()

			// Verify return value
			Expect(GetInstance()).To(Not(BeNil()))
		})
	})
})
