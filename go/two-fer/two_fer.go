/*
Package twofer is responsible for fair distribution of cookies and other treats.
*/
package twofer

import "fmt"

// ShareWith returns the string of the form "One for <name>, one for me."
// If the parameter name is an empty string, it is replaced with "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
