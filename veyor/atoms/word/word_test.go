package word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// success
	w, err := Parse("abc")
	assert.Equal(t, Word("abc"), w)
	assert.NoError(t, err)

	// error - invalid Word
	w, err = Parse("")
	assert.Zero(t, w)
	assert.EqualError(t, err, `invalid Word ""`)
}

func TestBool(t *testing.T) {
	// success - true
	ok := Word("abc").Bool()
	assert.True(t, ok)

	// success - false
	ok = Word("").Bool()
	assert.False(t, ok)
}

func TestNative(t *testing.T) {
	// success
	i := Word("abc").Native()
	assert.Equal(t, "abc", i)
}

func TestString(t *testing.T) {
	// success
	s := Word("abc").String()
	assert.Equal(t, "abc", s)
}
