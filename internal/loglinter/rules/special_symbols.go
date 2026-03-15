package rules

import (
	"unicode"
)

func isSpecialSymbol(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
		return false
	}
	return true
}

func CheckNoSpecialSymbols(msg string) bool {
	for _, ch := range msg {
		if isSpecialSymbol(ch) {
			return true
		}
	}
	return false
}
