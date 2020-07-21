package isogram

import "unicode"

// IsIsogram takes a string as input and return whether the input is isogram.
// definition of isogram is there are no duplicate letters.
func IsIsogram(input string) bool {
	set := make(map[rune]int)
	for _, val := range input {
		if !unicode.IsLetter(val) {
			continue
		}
		c := unicode.ToLower(val)
		if set[c] >= 1 {
			return false
		}
		set[c]++
	}
	return true
}
