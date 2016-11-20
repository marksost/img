// Helper functions for tests across the packages
package helpers

import (
	// Standard lib
	"fmt"
	"net/http"
	"net/http/httptest"
)

// GetMockServer returns a httptest server with the desired handler function
// based on the key passed in
func GetMockServer(key string) *httptest.Server {
	var handler http.Handler

	switch key {
	case "bad-request":
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Write headers and body
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{"code":400}`)
		})
	default:
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Write headers and body
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{"code":200,"foo":"bar","test":1234}`)
		})
	}

	return httptest.NewServer(handler)
}
