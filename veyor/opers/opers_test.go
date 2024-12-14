package opers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

func TestInit(t *testing.T) {
	// success
	assert.NotEmpty(t, len(Opers))
}

func TestWrap(t *testing.T) {
	// setup
	f := func(cs []cell.Cell) []cell.Cell { return []cell.Cell{cs[0] + cs[1]} }
	s := stack.New(1, 2)

	// success
	err := Wrap(2, f)(nil, s)
	assert.Equal(t, []cell.Cell{3}, s.Cells)
	assert.NoError(t, err)
}
