package luhn

import (
	"strings"
	"unicode"
)

var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Valid the input string with luhn algorithm
func Valid(input string) bool {
	inputTrimSpace := strings.Replace(input, " ", "", -1)
	if len(inputTrimSpace) < 2 {
		return false
	}
	odd := len(inputTrimSpace) & 1
	var sum int
	for i, c := range inputTrimSpace {
		if !unicode.IsDigit(c) {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}
