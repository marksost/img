// Test suite setup for the server package
package server

import (
	// Standard lib
	"io/ioutil"
	"testing"

	// Third-party
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Tests the server package
func TestConfig(t *testing.T) {
	// Register gomega fail handler
	RegisterFailHandler(Fail)

	// Have go's testing package run package specs
	RunSpecs(t, "Server Suite")
}

func init() {
	// Set logger output so as not to log during tests
	log.SetOutput(ioutil.Discard)
}
