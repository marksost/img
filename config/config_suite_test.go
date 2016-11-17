// Test suite setup for the config package
package config

import (
	// Standard lib
	"io/ioutil"
	"testing"

	// Third-party
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type (
	// Struct representing handleFields input data
	HandleFieldsTestData struct {
		Foo      string `env:"HANDLE_FIELDS_FOO"`
		FooEmpty string `env:""`
		Bar      int    `env:"HANDLE_FIELDS_BAR"`
		BarEmpty int    `env:""`
		Baz      bool   `env:"HANDLE_FIELDS_BAZ"`
		BazEmpty bool   `env:""`
		Test     struct {
			Foo string `env:"HANDLE_FIELDS_TEST_FOO"`
		}
	}
)

// Tests the config package
func TestConfig(t *testing.T) {
	// Register gomega fail handler
	RegisterFailHandler(Fail)

	// Have go's testing package run package specs
	RunSpecs(t, "Config Suite")
}

func init() {
	// Set logger output so as not to log during tests
	log.SetOutput(ioutil.Discard)
}
