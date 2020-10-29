package truncate

import (
	"unicode/utf8"
)

// WithoutFirstRune removes first rune from the given string
// Useful for removing # and @ from tags and ats strings
func WithoutFirstRune(str string) string {
	_, i := utf8.DecodeRuneInString(str)
	newStr := str[i:]
	return newStr
}
