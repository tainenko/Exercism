package anagram

import (
	"sort"
	"strings"
)

// Detect would check whether the string in candidates is anagram of subject and
// return all of the anagram candidates
func Detect(subject string, candidates []string) (res []string) {
	subject = strings.ToLower(subject)
	s := sortString(subject)
	for _, candidate := range candidates {
		c := strings.ToLower(candidate)
		if c == subject {
			continue
		}
		c = sortString(c)
		if c == s {
			res = append(res, candidate)
		}
	}
	return res
}

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
