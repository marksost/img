// Test suite setup for the operations package
package operations

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
	// Struct representing Dimension2Pixels input data
	Dimension2PixelsTestData struct {
		Dimension       string
		SourceDimension int64
		ReturnDimension int64
		ReturnWildcard  bool
		ReturnsError    bool
	}
	// Struct representing NewDimensionValues input data
	NewDimensionValuesTestData struct {
		Width        string
		Height       string
		SourceWidth  int64
		SourceHeight int64
		ReturnWidth  int64
		ReturnHeight int64
		ReturnsError bool
	}
	// Struct representing Ratio2Pixels input data
	Ratio2PixelsTestData struct {
		Dimension       string
		SourceDimension int64
		ReturnDimension int64
		ReturnsError    bool
	}
	// Struct representing RatioFromDimension input data
	RatioFromDimensionTestData struct {
		Numerator    int64
		Denominator  int64
		ReturnRatio  float64
		ReturnsError bool
	}
)

// Tests the operations package
func TestConfig(t *testing.T) {
	// Register gomega fail handler
	RegisterFailHandler(Fail)

	// Have go's testing package run package specs
	RunSpecs(t, "Image Operations Suite")
}

func init() {
	// Set logger output so as not to log during tests
	log.SetOutput(ioutil.Discard)
}
