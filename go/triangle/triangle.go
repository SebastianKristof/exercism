/*
Package triangle is for classifying triangles
*/
package triangle

import "math"

// Kind is a type representing triangle classes
// can take values _, NaT, Equ, Iso, Sca, Deg
type Kind int

// Possible kinds of triangles
const (
	_   Kind = iota
	NaT      // not a triangle
	Equ      // equilateral
	Iso      // isosceles
	Sca      // scalene
	Deg      // degenerate
)

// KindFromSides takes three side lengths
// and returns the kind of triangle (as type Kind)
func KindFromSides(a, b, c float64) Kind {
	// checking for wrong inputs
	for _, side := range []float64{a, b, c} {
		if side <= 0 || math.IsInf(side, 0) || math.IsNaN(side) {
			return NaT
		}
	}
	// another check for Not a triangle
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	// check for Degenerate
	if a+b == c || a+c == b || b+c == a {
		return Deg
	}
	// determine the triangle type
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
