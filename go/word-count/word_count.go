package wordcount

import(
	"strings"
	"regexp"
)

type Frequency map[string]int

func WordCount(phrase string)Frequency{

	frq := map[string]int{}

	formattedPhrase := formatString(phrase)


	frq = mapPhrase(formattedPhrase)

	return frq

}

func formatString(phrase string)[]string{
	reg := regexp.MustCompile("[^a-zA-Z0-9 ,']+")  
	processedPhrase1 := strings.ReplaceAll(phrase,","," ")
	processedPhrase2 := reg.ReplaceAllString(processedPhrase1, "")
	processedPhrase3 := strings.ToLower(processedPhrase2)
	processedPhrase4 := strings.Fields(processedPhrase3)

	return processedPhrase4

}

func mapPhrase(phrase []string)map[string]int{
	
	frequency := make(map[string]int)
	

	for _, words := range phrase{
		frequency[words]++
		}
	
	return frequency
}
