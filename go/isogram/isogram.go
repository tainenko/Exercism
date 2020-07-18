package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram takes a string as input and return whether the input is Isogram.
// definition of isogram is there are no duplicate letters.
func IsIsogram(input string) bool {
	set := make(map[rune]bool)
	for _, val := range strings.ToLower(input) {
		if !unicode.IsLetter(val) {
			continue
		}
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = true
	}
	return true
}
