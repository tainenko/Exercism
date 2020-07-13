// Package twofer
/*
Twofer package provide a function which return a formatted sentence.
*/
package twofer

import "fmt"

// Sharewith accepts a name variable of type string.
// Returns a string in the form "One for %s, One for me."
// If the name is an empty string, return "One for you, One for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
