package darts

import(
	"math"
)

//Pothag finds distance from center given 2 coordinates 
func Pothag(a float64, b float64) (c float64){
	c = math.Sqrt(a*a + b*b)
	return c
}

//Score finds the score of a dart thrown at a board based on the coordinates from center 
func Score (a float64, b float64) (score int){
	var dist float64

	dist = Pothag(a,b)

	switch {
	case dist <= 1:
		score = 10
	case dist <= 5:
		score = 5
	case dist <= 10:
		score = 1
	default:
		score = 0
	}

	return score
}
