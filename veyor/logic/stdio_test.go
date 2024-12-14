package logic

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

func TestInput(t *testing.T) {
	// setup
	s := stack.New()
	Stdin = bytes.NewBufferString("Test.\n")

	// success
	err := Input0(nil, s)
	assert.Equal(t, []cell.Cell{10, 46, 116, 115, 101, 84}, s.Cells)
	assert.NoError(t, err)
}

func TestPrint(t *testing.T) {
	// setup
	b := bytes.NewBuffer(nil)
	s := stack.New(0, 10, 46, 116, 115, 101, 84)
	Stdout = b

	// success
	err := PrintN(nil, s)
	assert.Equal(t, "Test.\n", b.String())
	assert.Equal(t, []cell.Cell{0}, s.Cells)
	assert.NoError(t, err)
}
