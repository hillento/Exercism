package isogram

import(
	"unicode"
	"strings"
)

//IsIsogram returns true if given input of S is an isogram
func IsIsogram (s string)bool{
	s = strings.ToLower(s)

	for i, j := range s {
		if unicode.IsLetter(j) && strings.ContainsRune(s[i+1:], j){
			return false
		}
	}
	return true
}
