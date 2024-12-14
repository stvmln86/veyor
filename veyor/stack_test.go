package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	// success
	s := NewStack(123)
	assert.Equal(t, []int{123}, s.Items)
}

func TestStackEmpty(t *testing.T) {
	// success - true
	ok := NewStack().Empty()
	assert.True(t, ok)

	// success - false
	ok = NewStack(123).Empty()
	assert.False(t, ok)
}

func TestStackLen(t *testing.T) {
	// success
	i := NewStack(123).Len()
	assert.Equal(t, 1, i)
}

func TestStackPeek(t *testing.T) {
	// success
	i := NewStack(123).Peek()
	assert.Equal(t, 123, i)
}

func TestStackPop(t *testing.T) {
	// setup
	s := NewStack(123)

	// success
	i := s.Pop()
	assert.Equal(t, 123, i)
	assert.Empty(t, s.Items)

	// error - Stack is empty
	defer func() { assert.Equal(t, "Stack is empty", recover()) }()
	s.Pop()
}

func TestStackPush(t *testing.T) {
	// setup
	s := NewStack()

	// success
	s.Push(123, 456)
	assert.Equal(t, []int{123, 456}, s.Items)
}

func TestStackString(t *testing.T) {
	// success
	s := NewStack(123, 456).String()
	assert.Equal(t, "123 456", s)
}
