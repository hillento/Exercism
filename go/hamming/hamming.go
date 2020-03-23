package hamming

import (
	"errors"
)

//Distance will return the number of differences in the 2 strings if they are the same length
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		myerror := errors.New("strings must be the same length")
		return 0, myerror
	}

	dist := 0
	ar, br := []rune(a), []rune(b)

	for i := range ar {
		if ar[i] != br[i] {
			dist++
		}
	}
	return dist, nil
}
