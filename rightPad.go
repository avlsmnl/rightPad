package rightPad

import (
	"strings"
	"unicode/utf8"
)

var default_char = ' '

// Format takes in a string and an int and returns the string
// Right-padded with spaces to the length of the int. If the
// string is already longer than the specified length, the
// original string is returned.
func Format(input string, length int) string {
	return FormatRune(input, length, default_char)
}

func FormatRune(input string, length int, r rune) string {
	currentSize := utf8.RuneCountInString(input)
	if currentSize >= length {
		return input
	}

	out := input + strings.Repeat(string(r), length-currentSize)
	return out
}
