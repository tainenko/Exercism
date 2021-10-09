package prime

// Factors compute the prime factors of a given natural number.
func Factors(n int64) []int64 {
	primes := []int64{}
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			primes = append(primes, i)
			n /= i
		}
	}
	return primes
}
