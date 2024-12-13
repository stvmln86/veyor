package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
)

func TestNew(t *testing.T) {
	// success
	s := New(123, 456)
	assert.Equal(t, []cell.Cell{123, 456}, s.Cells)
}

func TestEmpty(t *testing.T) {
	// success - true
	ok := New().Empty()
	assert.True(t, ok)

	// success - true
	ok = New(123, 456).Empty()
	assert.False(t, ok)
}

func TestLen(t *testing.T) {
	// success
	n := New(123, 456).Len()
	assert.Equal(t, 2, n)
}

func TestPop(t *testing.T) {
	// setup
	s := New(123)

	// success
	c, err := s.Pop()
	assert.Equal(t, cell.Cell(123), c)
	assert.NoError(t, err)

	// error - Stack is empty
	c, err = s.Pop()
	assert.Zero(t, c)
	assert.EqualError(t, err, `Stack is empty`)
}

func TestPopN(t *testing.T) {
	// setup
	s := New(123, 456)

	// success
	cs, err := s.PopN(2)
	assert.Equal(t, []cell.Cell{456, 123}, cs)
	assert.NoError(t, err)

	// error - Stack is insufficient
	cs, err = s.PopN(1)
	assert.Empty(t, cs)
	assert.EqualError(t, err, `Stack is insufficient`)
}

func TestPush(t *testing.T) {
	// setup
	s := New()

	// success
	s.Push(cell.Cell(123))
	assert.Equal(t, []cell.Cell{123}, s.Cells)
}

func TestPushAll(t *testing.T) {
	// setup
	s := New()

	// success
	s.PushAll([]cell.Cell{123, 456})
	assert.Equal(t, []cell.Cell{123, 456}, s.Cells)
}
