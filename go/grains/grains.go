package grains

import "errors"

// Square function calculate the number of grains of wheat on a chessboard given that the number
// on each square doubles.
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("square must be greater than zero or less than or equal to 64")
	}
	return uint64(1) << (n - 1), nil
}

// Total return the max value of uint64
func Total() uint64 {
	return ^uint64(0)
}
