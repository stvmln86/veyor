package opers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
)

func TestDupe1(t *testing.T) {
	// success
	cs := Dupe1([]cell.Cell{1})
	assert.Equal(t, []cell.Cell{1, 1}, cs)
}

func TestDrop1(t *testing.T) {
	// success
	cs := Drop1([]cell.Cell{1})
	assert.Empty(t, cs)
}

func TestRoll3(t *testing.T) {
	// success
	cs := Roll3([]cell.Cell{1, 2, 3})
	assert.Equal(t, []cell.Cell{2, 1, 3}, cs)
}

func TestSwap2(t *testing.T) {
	// success
	cs := Swap2([]cell.Cell{1, 2})
	assert.Equal(t, []cell.Cell{1, 2}, cs)
}
