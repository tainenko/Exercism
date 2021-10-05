package sieve

type number struct {
	value   int
	visited bool
}

// Sieve can find all the primes from 2 up to a given number.
func Sieve(limit int) []int {
	var primes []int
	
	var numbers = make([]number, limit+1)
	for i := 0; i <= limit; i++ {
		numbers[i].value = i
	}

	numbers[0].visited = true
	numbers[1].visited = true

	for _, num := range numbers {
		if num.visited {
			continue
		}
		primes = append(primes, num.value)
		for i := num.value; i <= limit; i += num.value {
			numbers[i].visited = true
		}
	}
	return primes
}
