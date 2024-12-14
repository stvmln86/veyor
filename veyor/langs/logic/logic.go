// Package logic implements parsing & evaluation functions.
package logic

import (
	"strings"

	"github.com/stvmln86/veyor/veyor/atoms/atom"
)

// Parse returns an Atom slice from a string.
func Parse(s string) ([]atom.Atom, error) {
	var as []atom.Atom

	for _, s := range strings.Fields(s) {
		a, err := atom.Atomise(s)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}
