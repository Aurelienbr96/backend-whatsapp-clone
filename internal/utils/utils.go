package utils

import (
	"regexp"
)

func RemoveWhiteSpace(s string) string {
	var nonNumericRegex = regexp.MustCompile(`[^\d+]`)

	return nonNumericRegex.ReplaceAllString(s, "")
}
