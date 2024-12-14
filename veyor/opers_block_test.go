package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpBreak(t *testing.T) {
	// success
	assertCode(t, "break")
	assert.True(t, Break)
}

func TestOpDef(t *testing.T) {
	// success
	assertCode(t, "def foo 123 end foo", 123)

	// error - "def" missing name/body
	defer func() { assert.Equal(t, `"def" missing name/body`, recover()) }()
	assertCode(t, "def end")

	// error - "def" name is wrong type
	defer func() { assert.Equal(t, `"def" name is wrong type`, recover()) }()
	assertCode(t, "def 123 456 end")
}

func TestOpIf(t *testing.T) {
	// success - true
	assertCode(t, "1 if 123 then", 123)

	// success - false
	assertCode(t, "0 if 123 then")
}

func TestOpLoopAndBreak(t *testing.T) {
	// success
	assertCode(t, "loop 123 break done", 123)
}
