// Package atom implements the Atom interface.
package atom

import "fmt"

// Atom is a single parsed program value.
type Atom interface {
	Bool() string
	Native() any
	String() string
}

// Atomise returns a parsed Atom from a string.
func Atomise(s string) (Atom, error) {
	switch {
	default:
		return nil, fmt.Errorf("invalid Atom %q", s)
	}
}
