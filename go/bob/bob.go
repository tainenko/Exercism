// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isStringUpper(input string) bool {
	r := []rune(input)

	for i := 0; i < len(r); i++ {
		ch := r[i]
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) == false {
				return false
			}
		}

	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if strings.Contains(remark, "?") && isStringUpper(remark) {
		return "Calm down, I know what I'm doing!"
	} else if strings.Contains(remark, "?") {
		return "Sure."
	} else if isStringUpper(remark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}
