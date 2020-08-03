package scrabble

import "strings"

//Score determines the point value a string would have in the game of scrabble
func Score(word string) int {
	word = strings.ToLower(word)
	var score int = 0

	for _, i := range word {
		switch i {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}

	}
	return score
}
