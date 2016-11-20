// downloader encapsulates all functionality around downloading an image from a URL
// including URL formation, HTTP requests, and data reading
package utils

import (
	// Standard lib
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// The default MIME type to return when no MIME type was set from the downloaded image
	DEFAULT_MIME_TYPE = "application/octet-stream"
)

type (
	// Struct representing a Downloader object used for downloading resources
	Downloader struct {
		data     []byte   // The raw data from the downloaded image
		mimeType string   // The detected MIME type of the downloaded image
		url      *url.URL // The URL to download the image from
	}
)

// NewDownloader creates a new `Downloader` and returns it
func NewDownloader(str string) *Downloader {
	// Create downloader instance
	d := &Downloader{}

	// Form URL instance from string and set downloader's URL
	d.url, _ = d.formUrl(str)

	// Return formed downloader
	return d
}

/* Begin main public functionality methods */

// Download makes an HTTP GET request for a given URL
// and returns the resulting data when possible
func (d *Downloader) Download() error {
	// Make HTTP GET request
	res, err := http.Get(d.url.String())
	if err != nil {
		return err
	}

	// Close body after processing
	defer res.Body.Close()

	// Check for success
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("URL returned status code other than 200: %d", res.StatusCode)
	}

	// Read data from response
	// TO-DO: Figure out how to test this...
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Set raw data
	d.data = data

	return nil
}

/* End main public functionality methods */

/* Begin internal propery methods */

// Data returns a byte slice representing the raw data from the downloaded image
func (d *Downloader) Data() []byte {
	return d.data
}

// MimeType returns a string representing the MIME type of the downloaded image
// NOTE: will return a default MIME type if none was previously set
func (d *Downloader) MimeType() string {
	// Check for empty MIME type and return default
	if d.mimeType == "" {
		return DEFAULT_MIME_TYPE
	}

	return d.mimeType
}

// Url returns a URL struct representing the parsed URL of the requested image
func (d *Downloader) Url() *url.URL {
	return d.url
}

/* End internal propery methods */

/* Begin utility methods */

// formUrl takes a URL string from a named request parameter, formats it,
// and returns it fully formed when possible
func (d *Downloader) formUrl(str string) (*url.URL, error) {
	// Remove leading and trailing whitespace and slashes
	str = strings.TrimLeft(str, " /")

	// Add scheme to URL
	str = "http://" + str

	// TO-DO: Handle URL forming, translation, etc

	return url.Parse(str)
}

/* End utility methods */
