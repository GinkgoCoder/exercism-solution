package wordy

import (
	"strconv"
	"strings"
	"unicode"
)

func Answer(question string) (int, bool) {
	runes := []rune(question)
	numStr, opStr := strings.Builder{}, strings.Builder{}
	res := 0
	for i, r := range runes {
		if unicode.IsDigit(r) || r == '+' || r == '-' {
			numStr.WriteRune(r)
		} else {
			if i > 0 && unicode.IsDigit(runes[i-1]) {
				val, err := strconv.Atoi(numStr.String())
				numStr.Reset()
				if err != nil {
					return -1, false
				}
				switch strings.TrimSpace(opStr.String()) {
				case "What is":
					res = val
				case "plus":
					res += val
				case "minus":
					res -= val
				case "multiplied by":
					res *= val
				case "divided by":
					res /= val
				default:
					return -1, false
				}
				opStr.Reset()
			}
			opStr.WriteRune(r)
		}
	}
	if opStr.String() != "?" {
		return -1, false
	}
	return res, true
}
