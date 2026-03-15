package rules

import (
	"unicode"
)

func CheckEnglish(msg string) bool {
	for _, ch := range msg {
		if unicode.IsLetter(ch) {
			if !unicode.Is(unicode.Latin, ch) {
				return true
			}
		}
	}
	return false
}
