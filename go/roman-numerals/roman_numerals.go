/*
Package romannumerals is responsible for roman numeral operations.
*/
package romannumerals

import (
	"errors"
	"strconv"
)

// ToRomanNumeral converts arabic numbers from 1 to 3000 to roman numbers.
func ToRomanNumeral(arabic int) (string, error) {

	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("unsupported number")
	}
	roman := ""
	arabicString := strconv.Itoa(arabic)
	place := len(arabicString)
	// digitsMap := buildDigitsMap() // enable for the 1st way
	for _, digitRune := range arabicString {
		if digitRune != '0' {

			// 1st way - 26 ms/op
			// romanRepresentation := digitsMap[place][digitRune]

			// 2nd way - 9 ms/op
			romanRepresentation := buildRomanDigitRepresentation(digitRune, place)

			roman += romanRepresentation
		}
		place--
	}
	return roman, nil
}

// this way is straightforward but slow
// lookups in 2d maps must be expensive
func buildDigitsMap() map[int]map[rune]string {
	return map[int]map[rune]string{
		1: map[rune]string{
			'1': "I",
			'2': "II",
			'3': "III",
			'4': "IV",
			'5': "V",
			'6': "VI",
			'7': "VII",
			'8': "VIII",
			'9': "IX",
		},
		2: map[rune]string{
			'1': "X",
			'2': "XX",
			'3': "XXX",
			'4': "XL",
			'5': "L",
			'6': "LX",
			'7': "LXX",
			'8': "LXXX",
			'9': "XC",
		},
		3: map[rune]string{
			'1': "C",
			'2': "CC",
			'3': "CCC",
			'4': "CD",
			'5': "D",
			'6': "DC",
			'7': "DCC",
			'8': "DCCC",
			'9': "CM",
		},
		4: map[rune]string{
			'1': "M",
			'2': "MM",
			'3': "MMM",
		},
	}
}

// this way is much faster
func buildRomanDigitRepresentation(digit rune, place int) string {
	symbols := map[int][]string{
		1: {"I", "V", "X"},
		2: {"X", "L", "C"},
		3: {"C", "D", "M"},
		4: {"M", "Unsupported", "Unsupported"},
	}

	symbolsToUse := symbols[place]
	s1, s2, s3 := symbolsToUse[0], symbolsToUse[1], symbolsToUse[2]

	r := ""
	switch digit {
	case '1':
		r = s1
	case '2':
		r = s1 + s1
	case '3':
		r = s1 + s1 + s1
	case '4':
		r = s1 + s2
	case '5':
		r = s2
	case '6':
		r = s2 + s1
	case '7':
		r = s2 + s1 + s1
	case '8':
		r = s2 + s1 + s1 + s1
	case '9':
		r = s1 + s3
	}
	return r
}
