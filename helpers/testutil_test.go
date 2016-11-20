// Tests the testutil.go file
package helpers

import (
	// Standard lib
	"io/ioutil"
	"net/http"
	"time"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("testutil.go", func() {
	Describe("`GetMockServer` method", func() {
		var (
			// Input for `GetMockServer` input
			input map[string]*GetMockServerTestData
		)

		Context("All non-timeout servers", func() {
			BeforeEach(func() {
				// Set input
				input = map[string]*GetMockServerTestData{
					"bad-request": &GetMockServerTestData{
						BodySubstring: "{\"code\":400}",
						StatusCode:    http.StatusBadRequest,
					},
					"default": &GetMockServerTestData{
						BodySubstring: "{\"code\":200,\"foo\":\"bar\",\"test\":1234}",
						StatusCode:    http.StatusOK,
					},
				}
			})

			It("Returns a mock server with predictable responses", func() {
				// Loop through test data
				for key, expected := range input {
					// Call method
					server := GetMockServer(key)

					// Make request against the server
					resp, err := http.Get(server.URL)

					// Read body
					body, _ := ioutil.ReadAll(resp.Body)

					// Verify return value
					Expect(err).To(Not(HaveOccurred()))
					Expect(string(body)).To(ContainSubstring(expected.BodySubstring))
					Expect(resp.StatusCode).To(Equal(expected.StatusCode))

					// Close body
					resp.Body.Close()
				}
			})
		})
	})
})
