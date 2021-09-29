package allyourbase

import (
	"errors"
	"math"
)

// ConvertToBase would convert the int array from inputBase to outputBase
func ConvertToBase(inputBase int, digit []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	num := 0
	var output []int
	for i := len(digit) - 1; i >= 0; i-- {
		if digit[i] < 0 || digit[i] >= inputBase {
			return output, errors.New("all digits must satisfy 0 <= d < input base")
		}
		order := int(math.Pow(float64(inputBase), float64(len(digit)-i-1)))
		num = digit[i]*order + num
	}
	if num == 0 {
		return []int{0}, nil
	}

	for num > 0 {
		output = append([]int{num % outputBase}, output...)
		num /= outputBase
	}
	return output, nil
}
