// serializers contains all custom serlializer functionality used throughout the application
package server

import (
	// Standard lib
	"fmt"

	// Internal
	"github.com/marksost/img/image/utils"

	// Third-party
	"github.com/kataras/go-serializer"
)

// setSerializers is used to set custom Iris serializers used throughout the application
func setSerializers() {
	// Loop through image MIME types, setting up a serializer for each
	for mime, _ := range utils.MimeTypes {
		server.UseSerializer(mime, serializer.SerializeFunc(imgSerializer))
	}
}

// imgSerializer is a serializer function used to properly serve images through Iris
// NOTE: `val` is expected to be a byte slice
// For more information on custom serializers, see
// https://docs.iris-go.com/serialize-engines.html
// https://github.com/kataras/go-serializer
func imgSerializer(val interface{}, options ...map[string]interface{}) ([]byte, error) {
	switch t := val.(type) {
	case []byte:
		return val.([]byte), nil
	default:
		return nil, fmt.Errorf("Invalid value type detected. Type was: %v", t)
	}
}
