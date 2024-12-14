package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertCode(t *testing.T, s string, is ...int) *Stack {
	st := NewStack()
	EvaluateString(Stlib+s, st)

	if len(is) == 0 {
		assert.Empty(t, st.Items)
	} else {
		assert.Equal(t, is, st.Items)
	}

	return st
}

func TestInit(t *testing.T) {
	// success
	assert.NotEmpty(t, Opers)
}

func TestEvaluate(t *testing.T) {
	// setup
	q := NewQueue(1, 2, "+")
	s := NewStack()

	// success
	Evaluate(q, s)
	assert.Equal(t, []int{3}, s.Items)
}

func TestEvaluateSlice(t *testing.T) {
	// setup
	s := NewStack()

	// success
	EvaluateSlice([]any{1, 2, "+"}, s)
	assert.Equal(t, []int{3}, s.Items)
}

func TestEvaluateString(t *testing.T) {
	// setup
	s := NewStack()

	// success
	EvaluateString("1 2 +", s)
	assert.Equal(t, []int{3}, s.Items)
}

func TestParse(t *testing.T) {
	// success
	as := Parse("123 abc\n")
	assert.Equal(t, []any{123, "abc"}, as)
}
