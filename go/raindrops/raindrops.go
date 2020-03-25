package raindrops

import "strconv"

//Converts number to sound based on its factors
//if not corresponding factors are found it returns the number as a string
func Convert(drop int) string {
	var snd string

	if drop%3 == 0 {
		snd += "Pling"
	}
	if drop%5 == 0 {
		snd += "Plang"
	}
	if drop%7 == 0 {
		snd += "Plong"
	}

	if snd == "" {
		return strconv.Itoa(drop)
	}

	return snd
}
