package rules

import (
	"unicode"
)

func CheckLowercase(msg string) bool {
	firstLetter := []rune(msg)[0]

	if unicode.IsUpper(firstLetter) {
		return true
	}

	return false
}
