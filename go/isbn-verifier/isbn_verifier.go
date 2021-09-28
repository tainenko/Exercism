package isbn

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[0-9]{9}[0-9X]$")

// IsValidISBN would check if the input string is valid.
func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if !re.MatchString(isbn) {
		return false
	}
	sum := 0
	for i := 0; i < 10; i++ {
		if isbn[i] == 'X' {
			sum += (10 - i) * 10
		} else {
			sum += (10 - i) * int(isbn[i]-'0')
		}
	}
	return sum%11 == 0
}
