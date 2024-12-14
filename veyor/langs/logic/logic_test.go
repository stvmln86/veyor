package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/atom"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
)

func TestParse(t *testing.T) {
	// success
	as, err := Parse("123 abc")
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, as)
	assert.NoError(t, err)
}
