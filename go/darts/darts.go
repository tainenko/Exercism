package darts

import "math"

// Score given a point in the target (defined by its real cartesian coordinates x and y),
// returns the correct amount earned by a dart landing in that point.
func Score(x, y float64) int {
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if dist <= 1 {
		return 10
	} else if dist <= 5 {
		return 5
	} else if dist <= 10 {
		return 1
	} else {
		return 0
	}
}
