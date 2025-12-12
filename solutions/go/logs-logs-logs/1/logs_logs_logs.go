package logs

import "strings"
import "math"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	emojis := [][]string{
        	{"🔍", "search"}, 
        	{"❗", "recommendation"}, 
        	{"☀", "weather"},
        }
    ans := "default"
    minIndex := math.MaxInt16 
    for _, emoji := range(emojis) {
        index := strings.Index(log, emoji[0]) 
        if index != -1 && index < minIndex {
            ans = emoji[1]
            minIndex = index
        }
    }
    return ans
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
