/*
Package dna includes functions for working with DNA strands.
*/
package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	// create a map so it can be written to
	h := make(map[rune]int)

	// initial zero values
	h['A'], h['T'], h['G'], h['C'] = 0, 0, 0, 0

	for i, nuc := range d {
		switch nuc {
		case 'A', 'T', 'G', 'C':
			h[nuc]++
		default:
			return h, fmt.Errorf("incorrect nucleotide %c at position %d", nuc, i)
		}
	}

	return h, nil
}
