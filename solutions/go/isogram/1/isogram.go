package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	set := map[rune]bool{}
	for _, rune := range word {
		rune = unicode.ToLower(rune)
		if unicode.IsLetter(rune) {
			if _, exists := set[rune]; exists {
				return false
			} else {
				set[rune] = true
			}
		}
	}
	return true
}
