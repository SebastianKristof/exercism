/*
Package isogram is for working with isograms.
*/
package isogram

import "unicode"

// IsIsogram determines if the input is an isogram.
// An isogram is a string in which no letter is seen more than once.
// It returns true if the input is an isogram, and false if it is not.
func IsIsogram(word string) bool {
	runeIsInWord := make(map[rune]bool)
	for _, r := range word {
		r = unicode.ToUpper(r)
		if unicode.IsLetter(r) && runeIsInWord[r] {
			return false
		}
		runeIsInWord[r] = true
	}
	return true
}
