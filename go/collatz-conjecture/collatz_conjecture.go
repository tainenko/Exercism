package collatzconjecture

import "errors"

// CollatzConjecture given a number n, return the number of steps required to reach 1.
func CollatzConjecture(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("n can't be zero")
	}
	if n < 0 {
		return 0, errors.New("n can't be negative")
	}
	step := 0
	for n != 1 {
		step++
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return step, nil
}
