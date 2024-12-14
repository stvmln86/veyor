package opers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
)

func TestAdd2(t *testing.T) {
	// success
	cs := Add2([]cell.Cell{1, 2})
	assert.Equal(t, []cell.Cell{3}, cs)
}

func TestDiv2(t *testing.T) {
	// success
	cs := Div2([]cell.Cell{6, 3})
	assert.Equal(t, []cell.Cell{2}, cs)
}

func TestMod2(t *testing.T) {
	// success
	cs := Mod2([]cell.Cell{5, 2})
	assert.Equal(t, []cell.Cell{1}, cs)
}

func TestMul2(t *testing.T) {
	// success
	cs := Mul2([]cell.Cell{3, 2})
	assert.Equal(t, []cell.Cell{6}, cs)
}

func TestSub2(t *testing.T) {
	// success
	cs := Sub2([]cell.Cell{5, 3})
	assert.Equal(t, []cell.Cell{2}, cs)
}
