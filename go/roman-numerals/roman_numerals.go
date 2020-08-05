package romannumerals

import "errors"

// translate type creates a structre for mapping digis to numerals
type translate struct {
	dig int
	rom string
}

//ToRomanNumeral converts numbers to roman numerals
func ToRomanNumeral(number int) (result string, err error) {
	//sends error if number is not within perameters of test.
	if number < 1 || number > 3000 {
		err = errors.New("Please enter number between 1 and 3000")
		return
	}

	//A map containing translate types. Used for iteration
	possible := []translate{
		translate{1000, "M"},
		translate{900, "CM"},
		translate{500, "D"},
		translate{400, "CD"},
		translate{100, "C"},
		translate{90, "XC"},
		translate{50, "L"},
		translate{40, "XL"},
		translate{10, "X"},
		translate{9, "IX"},
		translate{5, "V"},
		translate{4, "IV"},
		translate{1, "I"},
	}

	for _, i := range possible {
		for number >= i.dig {
			result += i.rom
			number -= i.dig
		}
	}

	return
}
