package opers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
)

func TestAdd2(t *testing.T) {
	// success
	c := Add2([]cell.Cell{1, 2})
	assert.Equal(t, cell.Cell(3), c)
}

func TestDiv2(t *testing.T) {
	// success
	c := Div2([]cell.Cell{6, 3})
	assert.Equal(t, cell.Cell(2), c)
}

func TestMod2(t *testing.T) {
	// success
	c := Mod2([]cell.Cell{5, 2})
	assert.Equal(t, cell.Cell(1), c)
}

func TestMul2(t *testing.T) {
	// success
	c := Mul2([]cell.Cell{3, 2})
	assert.Equal(t, cell.Cell(6), c)
}

func TestSub2(t *testing.T) {
	// success
	c := Sub2([]cell.Cell{5, 3})
	assert.Equal(t, cell.Cell(2), c)
}
