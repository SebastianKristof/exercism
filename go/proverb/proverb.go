/*
Package proverb includes functions for generating proverbs and other  pieces of wisdom.
*/
package proverb

import "fmt"

// Proverb takes a list of strings and returns a proverb following the sequence of causes and effects.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	result := []string{}
	rootCause := rhyme[0]
	if len(rhyme) > 1 {
		cause := rootCause
		for _, effect := range rhyme[1:] {
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", cause, effect))
			cause = effect
		}
	}
	result = append(result, fmt.Sprintf("And all for the want of a %s.", rootCause))

	return result
}
