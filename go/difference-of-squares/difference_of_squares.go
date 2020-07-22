package diffsquares

import "math"

// SquareOfSum takes a int as input and return the square of sum.
func SquareOfSum(target int) (result int) {
	for i := 1; i <= target; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares takes a int as input and return the sum of square.
func SumOfSquares(target int) (result int) {
	for i := 1; i <= target; i++ {
		result += i * i
	}
	return

}

// Difference function would return the absolute difference between SquareOfSum and SumOfSquares
func Difference(target int) int {
	return abs(SumOfSquares(target) - SquareOfSum(target))

}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
