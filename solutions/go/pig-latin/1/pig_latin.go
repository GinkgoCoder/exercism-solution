package piglatin

import "strings"

func Word(word string) string {
	if word == "" {
		return ""
	}

	runes := []rune(word)

	rule1 := func() bool {
		if runes[0] == 'a' || runes[0] == 'e' || runes[0] == 'i' || runes[0] == 'o' || runes[0] == 'u' {
			return true
		}
		if runes[0] == 'x' && runes[1] == 'r' {
			return true
		}
		if runes[0] == 'y' && runes[1] == 't' {
			return true
		}
		return false
	}
	rule3 := func() bool {
		if runes[0] == 'q' && runes[1] == 'u' {
			return true
		}
		return false
	}
	rule4 := func() bool {
		if runes[0] == 'y' {
			return true
		}
		return false
	}
	for i := range len(word) {
		switch {
		case rule1():
			return string(runes) + "ay"
		case rule3():
			return string(runes[2:]) + "quay"
		case i > 0 && rule4():
			return string(runes) + "ay"
		default:
			runes = append(runes[1:], runes[0])
		}
	}
	return word
}

func Sentence(sentence string) string {
	if sentence == "" {
		return ""
	}
	pigString := strings.Builder{}
	for _, word := range strings.Fields(sentence) {
		pigString.WriteString(Word(word))
		pigString.WriteString(" ")
	}
	return strings.TrimSpace(pigString.String())
}
