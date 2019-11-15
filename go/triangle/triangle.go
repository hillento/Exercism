package triangle

import "math"

//Kind of triangle
type Kind int

const (
	// NaT = not a triangle
	NaT Kind = iota
	//Equ = equilateral
	Equ
	//Iso = isosceles
	Iso
	// Sca =scalene
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if Valid(a, b, c) && Valid(b, c, a) && Valid(c, a, b) {
		if a == b && b == c {
			return Equ
		}
		if a == b || a == c || b == c {
			return Iso
		}
		return Sca
	}
	return NaT
}

//Valid determines if the leg lengths can make a valid triangle
func Valid(a, b, c float64) bool {
	return !math.IsNaN(a) && !math.IsInf(a, 0) && a > 0 && a <= b+c
}
