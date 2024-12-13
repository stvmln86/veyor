// Package word implements the Word Atom type.
package word

import (
	"fmt"
	"strings"
)

// Word is a parsed reference value.
type Word string

// Parse returns a parsed Word from a string.
func Parse(s string) (Word, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("invalid Word %q", s)
	}

	return Word(s), nil
}

// Bool returns the Word as a boolean.
func (w Word) Bool() bool {
	return string(w) != ""
}

// Native returns the Word as a native value.
func (w Word) Native() any {
	return string(w)
}

// String returns the Word as a string.
func (w Word) String() string {
	return string(w)
}
