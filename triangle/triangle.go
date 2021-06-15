// Package triangle provides some utilities for triangles.
package triangle

import (
	"math"
)

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// All sides have to be of length > 0 and the sum of the lengths of any two sides
// must be greater than or equal to the length of the third side.
func isTriangle(in ...float64) bool {
	if len(in) == 3 {
		for _, v := range in {
			if math.IsInf(v, 0) {
				return false
			} else if math.IsNaN(v) {
				return false
			} else if v <= 0 {
				return false
			}
		}
		if in[0]+in[1] >= in[2] && in[0]+in[2] >= in[1] && in[1]+in[2] >= in[0] {
			return true
		}
	} else {
		return false
	}
	return false
}

// KindFromSides checks the kind of the triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT
	if isTriangle(a, b, c) {
		if a == b && b == c { // an equilateral triangle has all three sides the same length
			k = Equ
		} else if a == b || a == c || b == c { // an isosceles triangle has at least two sides the same length
			k = Iso
		} else if a != b && a != c && b != c { // a scalene triangle has all sides of different lengths
			k = Sca
		}

	}
	return k
}
