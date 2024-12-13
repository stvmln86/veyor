// Package atom implements the Atom interface.
package atom

import (
	"fmt"

	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
)

// Atom is a single parsed program value.
type Atom interface {
	// Bool returns the Atom as a boolean.
	Bool() bool

	// Native returns the Atom as a native value.
	Native() any

	// String returns the Atom as a string.
	String() string
}

// Atomise returns a parsed Atom from a string.
func Atomise(s string) (Atom, error) {
	if c, err := cell.Parse(s); err == nil {
		return c, nil
	}

	if w, err := word.Parse(s); err == nil {
		return w, nil
	}

	return nil, fmt.Errorf("invalid Atom %q", s)
}
