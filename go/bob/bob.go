// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

func hasAlpha(s string) bool {
	reg := regexp.MustCompile(`[a-zA-Z]+`)
	return reg.MatchString(s)
}

func isYelling(s string) bool {
	if !hasAlpha(s) {
		return false
	}
	upper := strings.ToUpper(s)
	return upper == s
}

func isQuestion(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[len(s)-1] == '?'
}

func isYellingQustion(s string) bool {
	return isQuestion(s) && isYelling(s) && hasAlpha(s)
}

func isSilence(s string) bool {
	return len(s) == 0
}

// Hey responses to the question with limited answers.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isSilence(remark) {
		return "Fine. Be that way!"
	}
	if isYellingQustion(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion(remark) {
		return "Sure."
	}
	if isYelling(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
