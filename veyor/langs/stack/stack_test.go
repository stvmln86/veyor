package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	stack := New(0)
	assert.Equal(t, []int{0}, stack.Items)
}

func TestEmpty(t *testing.T) {
	// success - true
	ok := New().Empty()
	assert.True(t, ok)

	// success - false
	ok = New(0).Empty()
	assert.False(t, ok)
}

func TestLen(t *testing.T) {
	// success
	i := New(0).Len()
	assert.Equal(t, 1, i)
}

func TestPop(t *testing.T) {
	// setup
	stack := New(0)

	// success
	i, err := stack.Pop()
	assert.Equal(t, 0, i)
	assert.Empty(t, stack.Items)
	assert.NoError(t, err)

	// error - Stack is empty
	i, err = stack.Pop()
	assert.Zero(t, i)
	assert.EqualError(t, err, "Stack is empty")
}

func TestPush(t *testing.T) {
	// setup
	stack := New()

	// success
	stack.Push(0)
	assert.Equal(t, []int{0}, stack.Items)
}

func TestString(t *testing.T) {
	// success
	s := New(0, 1, 2).String()
	assert.Equal(t, "0 1 2", s)
}
