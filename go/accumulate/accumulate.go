/*
Package accumulate introduces the students to functional programming
*/
package accumulate

// Accumulate applies the operation to each element of the collection,
// and returns the resulting collection
func Accumulate(collection []string, operation func(string) string) []string {
	// pre-allocating a slice for efficiency, so we don't have to append
	result := make([]string, len(collection))

	for i, el := range collection {
		result[i] = operation(el)
	}

	return result
}
