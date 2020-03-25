package strand

// ToRNA takes converts a string to a rune and converts each letter based on the 'transcription' map
// ToRNA returns the converted runes as a string
func ToRNA(dna string) string {

	rdna := []rune(dna)
	rna := ""

	transcription := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for i := range rdna {
		dex := rdna[i]
		rna += string(transcription[dex])
	}

	return rna

}
