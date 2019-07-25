/*
Package strain is an ugly way to implement Keep and Discard methods for collections of integers,
lists, and strings. Some methods are not implemented yet.
It may be rewritten later with a unified API using interfaces.
*/
package strain

// Ints is a collection of integers.
type Ints []int

// Lists is a collection of slices.
type Lists [][]int

// Strings is a collection of strings.
type Strings []string

//Keep keeps those items from the slice of ints which satisfy the predicate.
func (ints Ints) Keep(f func(int) bool) Ints {
	var result Ints
	for _, int := range ints {
		if f(int) {
			result = append(result, int)
		}
	}
	return result
}

//Discard discards the items from the slice of ints which do not satisfy the predicate.
func (ints Ints) Discard(f func(int) bool) Ints {
	var result Ints
	for _, int := range ints {
		if !f(int) {
			result = append(result, int)
		}
	}
	return result
}

// Keep keeps those items from the slice of slices which satisfy the predicate.
func (lists Lists) Keep(f func([]int) bool) Lists {
	var result Lists
	for _, list := range lists {
		if f(list) {
			result = append(result, list)
		}
	}
	return result
}

// Keep keeps those items from the slice of strings which satisfy the predicate.
func (strs Strings) Keep(f func(string) bool) Strings {
	var result Strings
	for _, str := range strs {
		if f(str) {
			result = append(result, str)
		}
	}
	return result
}
