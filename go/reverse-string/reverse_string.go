/*
Package reverse includes functions for reversing strings.
*/
package reverse

// Reverse will take a string and return it in the reversed order.
func Reverse(s string) string {
	source := []rune(s)
	sourceLength := len(source)
	if sourceLength == 0 {
		return ""
	}

	reverse := make([]rune, sourceLength)
	for i := range source {
		reverse[i] = source[sourceLength-1-i]
	}
	return string(reverse)
}
