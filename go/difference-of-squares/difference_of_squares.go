package diffsquares

// SquareOfSum takes a int as input and return the square of sum.
// formula: 1+2+3+...+n=n*(n+1)/2
func SquareOfSum(n int) (result int) {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares takes a int as input and return the sum of square.
// formula: 1*1+2*2+3*3+...+n*n=n*(n+1)*(2n+1)/6
func SumOfSquares(n int) (result int) {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference function would return the  difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
