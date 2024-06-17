package gohelpers

import (
	"fmt"
	"strconv"
)

func StringToFloat64(stringValue string) float64 {
	// Convert string value to float64
	float64res, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		fmt.Println("goHelpers: Error converting string to float64:", err)
	}

	return float64res
}
