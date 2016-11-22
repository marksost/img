// converters are used to convert values from one type to another
package helpers

import (
	// Standard lib
	"strconv"

	// Third-party
	log "github.com/Sirupsen/logrus"
)

// Bool2String converts a bool to a string
func Bool2String(v bool) string {
	return strconv.FormatBool(v)
}

// Float642String converts a float64 to a string
func Float642String(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// Int2String converts an int to a string
func Int2String(v int) string {
	return strconv.Itoa(v)
}

// Int642String converts an int64 to a string
func Int642String(v int64) string {
	return strconv.Itoa(int(v))
}

// Interface2String attempts to determine the underlying type of an interface and returns it as a string
func Interface2String(i interface{}) string {
	// Attempt to cast attribute based on it's underlying type
	switch t := i.(type) {
	case float64:
		return Float642String(i.(float64))
	case int:
		return Int2String(i.(int))
	case int64:
		return Int642String(i.(int64))
	case string:
		return i.(string)
	default:
		// Log unsupported type
		log.WithField("type", t).Warn("Interface is of unsupported type")

		return ""
	}
}

// MapFromInterface type-asserts the raw API response interface as a map
// so that other methods can more-easily access it's properties
func MapFromInterface(i interface{}) map[string]interface{} {
	return i.(map[string]interface{})
}

// String2Int converts a string to an int
func String2Int(v string) int {
	return int(String2Int64(v))
}

// String2Int64 converts a string to an int64
func String2Int64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		// Log conversion error
		log.WithFields(log.Fields{
			"string": v,
			"error":  err.Error(),
		}).Warn("Error converting string to int64")

		return 0
	}

	return i
}
