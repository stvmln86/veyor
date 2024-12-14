package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertCode(t *testing.T, s string, is ...int) *Stack {
	st := NewStack()
	EvaluateString(s, st)
	assert.Equal(t, is, st.Items)
	return st
}

func TestInit(t *testing.T) {
	// success
	assert.NotEmpty(t, Opers)
}
