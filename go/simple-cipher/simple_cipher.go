package cipher

import (
	"regexp"
	"strings"
)

// Vigenere specify a key and use that for the shift distance.
type Vigenere struct {
	key string
}

var numOfLetters int32 = 26
var re = regexp.MustCompile("[^a-zA-Z]")

// Encode return the cipher
func (cipher *Vigenere) Encode(s string) string {
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	var builder strings.Builder
	for i, v := range s {
		c := rune(cipher.key[i%len(cipher.key)])
		builder.WriteRune('a' + ((v-'a')+(c-'a'))%numOfLetters)
	}
	return builder.String()
}

// Decode return the plain text
func (cipher *Vigenere) Decode(s string) string {
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	var builder strings.Builder
	for i, v := range s {
		c := rune(cipher.key[i%len(cipher.key)])
		builder.WriteRune('a' + ((v-'a')-(c-'a')+numOfLetters)%numOfLetters)
	}
	return builder.String()
}

// NewVigenere return a Vigenere  instance
func NewVigenere(key string) Cipher {
	min := 'z'
	max := 'a'
	for _, c := range key {
		if c < min {
			min = c
		}

		if c > max {
			max = c
		}
	}
	if max > 'z' || min < 'a' || max == 'a' {
		return nil
	}
	return &Vigenere{key: key}
}

// NewShift returns a shift cipher which will shift pain text by k.
func NewShift(k int) Cipher {
	if k == 0 || k > 25 || k < -25 {
		return nil
	}
	return NewVigenere(string('a' + (rune(k)+numOfLetters)%numOfLetters))
}

// NewCaesar return a simple shift cipher which shift 3
func NewCaesar() Cipher {
	return NewShift(3)
}
