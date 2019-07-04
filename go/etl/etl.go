/*
Package etl is a helper package for transforming
the letter scores from the old format to the new format.
*/
package etl

import "strings"

// Transform extracts letter scores from the old format and loads them into the new format.
func Transform(oldScores map[int][]string) map[string]int {

	newScores := make(map[string]int)

	for score, letters := range oldScores {
		for _, letter := range letters {
			newScores[strings.ToLower(letter)] = score
		}
	}

	return newScores
}
