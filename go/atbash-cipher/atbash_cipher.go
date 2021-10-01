package atbash

import (
	"regexp"
	"strconv"
	"strings"
)

// Atbash encrypt the original string and return
func Atbash(plain string) string {
	var cipher []rune
	for _, r := range plain {

		if r >= 'a' && r <= 'z' {
			r = rune(25 - int(r) + 'a' + 'a')
		} else if r >= 'A' && r <= 'Z' {
			r = rune(25 - int(r) + 'A' + 'a')
		} else if r >= '0' && r <= '9' {
		} else {
			continue
		}
		cipher = append(cipher, r)
	}

	return strings.TrimSpace(insertNewLines(string(cipher), 5))
}

func insertNewLines(s string, n int) string {
	var r = regexp.MustCompile("(.{" + strconv.Itoa(n) + "})")
	return r.ReplaceAllString(s, "$1 ")
}
