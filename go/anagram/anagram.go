package anagram

import(
	"strings"
	"sort"
)

//Detect is used to return anagrams of a word from a list of candidates
func Detect(subject string, candidates []string )(anagrams []string){
	
	lowerSubject := strings.ToLower(subject)

	for _, potential := range candidates{
		lowerCandidte := strings.ToLower(potential)

		if lowerCandidte != lowerSubject && matchingLetters(lowerSubject, lowerCandidte){
			anagrams = append(anagrams, potential)
		}
	}
	return anagrams
}
//matchingLetters returns true if the two strings passed in have the same letters
func matchingLetters(subject string, candidate string)bool{
	subject = orderString(subject)
	candidate = orderString(candidate)

	if subject == candidate{
		return true
	}
	return false
}
//orderString puts the passed in string in alphabetical order
func orderString(word string)string{
	s := strings.Split(word, "")
	sort.Strings(s)

	sortedString := strings.Join(s,"")

	return sortedString
}


