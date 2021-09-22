package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is type of map
type Frequency map[string]int

var re = regexp.MustCompile(`\w+('\w)*`)

// WordCount will return the frequency of each word in phrase
func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	f := make(Frequency)
	for _, word := range re.FindAllString(phrase, -1) {
		f[word]++
	}
	return f
}
