package encode

import (
	"strconv"
	"strings"
	"unicode"
)

type Buffer struct {
	strings.Builder
}

func (b *Buffer) Encode(r rune, count int) {
	if count == 0 {
		return
	}
	if count > 1 {
		b.WriteString(strconv.Itoa(count))
	}
	b.WriteRune(r)
}

func (b *Buffer) Decode(r rune, count int) {
	for i := 0; i < count; i++ {
		b.WriteRune(r)
	}
}

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}
	res := Buffer{}
	prev, count := '0', 0
	for _, r := range input {
		if r != prev {
			res.Encode(prev, count)
			prev = r
			count = 1
			continue
		}
		count++
	}
	res.Encode(prev, count)
	return res.String()
}

func RunLengthDecode(input string) string {
	res := Buffer{}
	count := 0
	for _, r := range input {
		if unicode.IsDigit(r) {
			count = count*10 + int(r-'0')
		} else {
			count = max(count, 1)
			res.Decode(r, count)
			count = 0
		}
	}
	return res.String()
}
