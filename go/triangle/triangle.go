// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Kind represents the type of triangle
type Kind int

const (
	// NaT not a triangle
	NaT = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	sides := [3]float64{a, b, c}
	sort.Float64s(sides[:])
	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}

	a, b, c = sides[0], sides[1], sides[2]
	if a+b < c {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
