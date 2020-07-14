// Package hamming provide a method to calculate the Hamming Distance
package hamming

import "errors"

// Distance function return a humming distance of two DNA string.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("the Hamming distance is only defined for sequences of equal length")
	}
	var count int
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
