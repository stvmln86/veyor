// Package cell implements the Cell Atom type.
package cell

import (
	"fmt"
	"strconv"
)

// Cell is a parsed integer value.
type Cell int64

// Parse returns a parsed Cell from a string.
func Parse(s string) (Cell, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid Cell %q", s)
	}

	return Cell(i), nil
}

// Bool returns the Cell as a boolean.
func (c Cell) Bool() bool {
	return int64(c) != 0
}

// Native returns the Cell as a native value.
func (c Cell) Native() any {
	return int64(c)
}

// String returns the Cell as a string.
func (c Cell) String() string {
	return strconv.FormatInt(int64(c), 10)
}
