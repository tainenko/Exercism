package prime

import "math"

// Nth returns nth prime
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	cnt := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			cnt++
		}
		if cnt == n {
			return i, true
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
