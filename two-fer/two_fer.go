// Package twofer implements a simple library for crating a twofer sentence.
package twofer

import "fmt"

// ShareWith returns a string which contains the given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
