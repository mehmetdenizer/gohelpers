package gohelpers

// Filters numbers as digits, removes spaces and all non-numeric characters

import (
	"regexp"
	"strings"
)

func JustDigits(input string) string {
	// Create a regex pattern that matches only numbers
	re := regexp.MustCompile(`\d+`)
	// Find all matches
	digits := re.FindAllString(input, -1)
	// Combine the found numbers
	return strings.Join(digits, "")
}
