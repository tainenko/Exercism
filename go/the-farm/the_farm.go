package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		fodder *= 2
	} else if err != nil {
		return 0, err
	}

	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows < 0 {
		return 0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	}

	return fodder / float64(cows), nil
}
