package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStlib(t *testing.T) {
	// setup
	b := mockStreams("test\n")

	// success - boolean functions
	assertCode(t, "2 not", -2)

	// success - conditional functions
	assertCode(t, "1 even? 2 even?", 0, 1)
	assertCode(t, "1 odd? 2 odd?", 1, 0)
	assertCode(t, "0 zero? 1 zero?", 1, 0)

	// success - miscellaneous functions
	assertCode(t, "·")

	// success - stack functions
	assertCode(t, "1 2 3 clear")
	assertCode(t, "1 2 drop", 1)

	// success - standard i/o functions
	assertCode(t, "0 10 116 115 101 116 print0")
	assertCode(t, "input", 0, 10, 116, 115, 101, 116)
	assert.Equal(t, "test\n", b.String())
}
