package armstrong

import (
	"fmt"
	"math"
)

// IsNumber returns true if input number is An Armstrong. Armstrong number is a number that is the sum of its own digits
// each raised to the power of the number of digits.
func IsNumber(number int) bool {
	if number == 0 || number <= 9 {
		return true
	}
	s := fmt.Sprint(number)
	digit := len(s)
	var total float64
	for _, v := range s {

		total += math.Pow(float64(v-'0'), float64(digit))
	}
	return number == int(total)
}
