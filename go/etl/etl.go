package etl

import(
	"strings"
)

//Transform returns a map with a 1 to 1 score to letter assignment
func Transform(og map[int] []string) map[string]int{
	
	transformed := make(map[string]int)

	for score, letters := range og {
		for _, letter := range letters{
			transformed[strings.ToLower(letter)] = score
		}
	}
	return transformed
}