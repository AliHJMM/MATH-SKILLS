package functions

import "regexp"

// isValidNumber checks if the input string represents a valid number (no non-numeric characters).
func isValidNumber(s string) bool {
	match, _ := regexp.MatchString(`^\d+$`, s)
	return match
}
