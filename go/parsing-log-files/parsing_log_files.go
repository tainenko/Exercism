package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// IsValidLine returns false if a string is not valid otherwise true.
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|ERR|FTL)\]+`)
	return re.MatchString(text)
}

// SplitLogLine takes a line and returns an array of strings each of which contains a field.
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords provides an indication of the likely scale of the manual exercise.
func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompilePOSIX(`"[a-zA-Z ]*pass[wW]ord[a-zA-Z ]*"`)
	cnt := 0
	for _, line := range lines {
		if re.MatchString(line) {
			cnt += 1
		}
	}
	return cnt
}

// RemoveEndOfLineText removes the end of line string from text
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of[\w-]*`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName processes log lines
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		if matches := re.FindStringSubmatch(line); len(matches) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", matches[1], lines[i])
		}
	}
	return lines
}
