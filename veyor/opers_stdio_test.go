package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpEval(t *testing.T) {
	// success
	assertCode(t, "0 0 43 32 50 32 49 eval", 0, 3)
}

func TestOpInput(t *testing.T) {
	// setup
	mockStreams("test\n")

	// success
	assertCode(t, "input", 0, 10, 116, 115, 101, 116)
}

func TestOpPrint(t *testing.T) {
	// setup
	b := mockStreams("")

	// success
	assertCode(t, "116 print")
	assert.Equal(t, "t", b.String())
}
