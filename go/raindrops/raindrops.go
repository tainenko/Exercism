// Package raindrops provide a Convert function
// which receive a integer and return a string
package raindrops

import "strconv"

// Convert receive a integer.
// if the number is not a multiple of 3 or 5 or 7, return the string of number.
// if the number has [3, 5, 7] as a factor, add ['Pling', 'Plang', 'Plong'] to the result
func Convert(num int) (result string) {
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(num)
	}
	return
}
