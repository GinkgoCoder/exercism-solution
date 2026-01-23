package cryptosquare

import (
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

func Encode(pt string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	text := re.ReplaceAllString(pt, "")
	size := utf8.RuneCountInString(text)
	r, c := int(math.Ceil(math.Sqrt(float64(size)))), int(math.Floor(math.Sqrt(float64(size))))
	if c*r < size {
		c++
	}
	runes := make([][]rune, r)
	index := 0
	for _, c := range text {
		runes[index%r] = append(runes[index%r], c)
		index++
	}
	res := []string{}
	for _, row := range runes {
		res = append(res, strings.ToLower(string(row)+strings.Repeat(" ", c-len(row))))
	}
	return strings.Join(res, " ")
}
