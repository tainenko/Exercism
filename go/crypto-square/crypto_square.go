package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`[^a-zA-Z0-9]`)

// Encode returns the encoded version of input english text
func Encode(input string) string {
	normalized := reg.ReplaceAllString(input, "")
	normalized = strings.ToLower(normalized)
	c := int(math.Ceil(math.Sqrt(float64(len(normalized)))))
	r := c - 1
	if c*r < len(normalized) {
		r = c
	}
	res := make([]string, c)
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			idx := i + c*j
			if idx < len(normalized) {
				res[i] += string(normalized[idx])
			} else {
				res[i] += " "
			}
		}
	}
	return strings.Join(res, " ")
}
