package hamming

import (
	"errors"
)

//Distance will return the number of differences in the 2 strings if they are the same length
func Distance(a, b string) (int, error) {
	dist := 0
	strandALen := len(a)
	strandBLen := len(b)

	if strandALen != strandBLen {
		myerror := errors.New("strings must be the same length")
		return -1, myerror
	}

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
