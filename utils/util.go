package utils

import (
	"strconv"
)

// StringToUint converts a string to a uint.
// If the conversion fails, it returns 0.
func StringToUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		// Return 0 if there's an error
		return 0
	}
	return uint(val)
}
