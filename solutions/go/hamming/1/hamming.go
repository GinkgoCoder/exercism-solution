package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("two strings should have the same length")
	}
	res := 0
	brunes := []rune(b)
	for i, c := range a {
		if brunes[i] != c {
			res += 1
		}
	}
	return res, nil
}
