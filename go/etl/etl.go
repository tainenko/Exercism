package etl

import "strings"

// Transform the legacy data format to the shiny new format
// The old system stored a list of letters per score
// The shiny system stores a list of score per the letters in lower-case regardless of the case of the input letters.
func Transform(m map[int][]string) map[string]int {
	var res = make(map[string]int)
	for k, v := range m {
		for _, s := range v {
			res[strings.ToLower(s)] = k
		}
	}
	return res
}
