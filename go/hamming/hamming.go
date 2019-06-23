/*
Package hamming includes functions to measure the process of evolution
for those users that happen to believe in it.
*/
package hamming

import (
	"errors"
)

// Distance is a function that calculates the Hamming distance between two strings.
func Distance(a, b string) (int, error) {

	// converting to slices of runes, because incoming strings
	// might include characters encoded by more than one byte
	ra, rb := []rune(a), []rune(b)

	if len(ra) != len(rb) {
		return 0, errors.New("strings must be of equal length")
	}

	distance := 0
	for i := range ra {
		if ra[i] != rb[i] {
			distance++
		}
	}

	return distance, nil
}
