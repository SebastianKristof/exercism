// Package strand includes functions to analyze strands of nucleotides
// and perform operations on them
package strand

import (
	"regexp"
	"strings"
)

// ToRNA transcribes a DNA sequence to its complement RNA sequence
func ToRNA(dna string) string {

	// 1. check that input is a correct DNA string
	dna = strings.ToUpper(dna)
	extraChars, _ := regexp.MatchString("[^ACGT]", dna)
	if extraChars {
		return "error: wrong characters in the DNA sequence"
	}

	// 2. define complement nucleotides
	// func transcribeToComplement below

	// 3. perform transcription
	runesDNA := []rune(dna)
	runesRNA := make([]rune, len(runesDNA))
	for i := range runesDNA {
		runesRNA[i] = transcribeToComplement(runesDNA[i])
	}

	return string(runesRNA)
}

// transcribeToComplement assumes that the input is already sanitized
func transcribeToComplement(nucDNA rune) (nucRNA rune) {
	switch nucDNA {
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	}
	return
}
