package pangram

import "unicode"

// IsPangram return true if input string contains every letter of the alphabet
func IsPangram(s string) bool {
	var m = map[rune]int{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			m[unicode.ToLower(r)]++
		}
	}
	return len(m) == 26
}
