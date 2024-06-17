package gohelpers

import (
	"fmt"
	"strings"
)

// A helper function to separate with commas

func FormatFloatWithComma(value float64) string {
	strValue := fmt.Sprintf("%.2f", value)
	return strings.Replace(strValue, ".", ",", 1)
}
