package raindrops

import "strconv"

//Convert takes and integer and returns a string that is dependant on the number's factors.
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
