package grains

import (
	"errors"
	"math"
)

// Square function calculate the number of grains of wheat on a chessboard given that the number
// on each square doubles.
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("square must be greater than zero or less than or equal to 64")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

// Total return the value of 1+2+4+8+...+2**64
func Total() uint64 {
	// formula: 1+2+4+8+...+2**64=2**65-1
	return 2*uint64(math.Pow(2, float64(64))) - 1
}
