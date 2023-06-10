package utils

import "unicode"

func CapitalizeFirstChar(str string) string {
	// Convert the string to a rune slice
	runes := []rune(str)

	// Check if the string is not empty
	if len(runes) > 0 {
		// Capitalize the first rune
		runes[0] = unicode.ToUpper(runes[0])
	}

	// Convert the rune slice back to a string
	return string(runes)
}
