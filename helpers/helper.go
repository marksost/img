// helpers package defines various helper functions for use throughout the application
package helpers

// SliceContains returns true if a slice of strings includes a specific string
func SliceContains(needle string, haystack []string) bool {
	for _, value := range haystack {
		if needle == value {
			return true
		}
	}

	return false
}
