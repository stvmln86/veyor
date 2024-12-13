package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	// success - true
	ok := Is("123")
	assert.True(t, ok)

	// success - false
	ok = Is("nope")
	assert.False(t, ok)
}

func TestParse(t *testing.T) {
	// success
	c, err := Parse("123")
	assert.Equal(t, Cell(123), c)
	assert.NoError(t, err)

	// error - invalid Cell
	c, err = Parse("")
	assert.Zero(t, c)
	assert.EqualError(t, err, `invalid Cell ""`)
}

func TestBool(t *testing.T) {
	// success - true
	ok := Cell(123).Bool()
	assert.True(t, ok)

	// success - false
	ok = Cell(0).Bool()
	assert.False(t, ok)
}

func TestNative(t *testing.T) {
	// success
	i := Cell(123).Native()
	assert.Equal(t, int64(123), i)
}

func TestString(t *testing.T) {
	// success
	s := Cell(123).String()
	assert.Equal(t, "123", s)
}
