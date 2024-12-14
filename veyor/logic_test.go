package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	// setup
	q := NewQueue(1, 2, "+")
	s := NewStack()

	// success
	Evaluate(q, s)
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
