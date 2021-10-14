package savings

import (
	"math"
)

// InterestRate calculates the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return -3.213
	}
	if balance < 1000 {
		return 0.5
	}
	if balance < 5000 {
		return 1.621
	}
	return 2.475
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * math.Abs(float64(InterestRate(balance))) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	for i := 1; ; i++ {
		balance = AnnualBalanceUpdate(balance)
		if balance > targetBalance {
			return i
		}
	}
}
