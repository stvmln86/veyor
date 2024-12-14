// Package parse implements string parsing functions.
package parse

import (
	"strings"

	"github.com/stvmln86/veyor/veyor/atoms/atom"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
)

// Parse returns an Atom slice from a string.
func Parse(s string) ([]atom.Atom, error) {
	var as []atom.Atom
	var rs []rune
	var q bool

	for _, r := range strings.TrimSpace(s) {
		if q || r != ' ' {
			rs = append(rs, r)
		}

		if r == '"' {
			if q = !q; !q {
				for i := len(rs) - 2; i > 0; i-- {
					as = append(as, cell.Cell(rs[i]))
				}

				rs = []rune{}
			}
		}

		if !q && r == ' ' && len(rs) > 0 {
			a, err := atom.Atomise(string(rs))
			if err != nil {
				return nil, err
			}

			as = append(as, a)
			rs = []rune{}
		}
	}

	return as, nil
}
