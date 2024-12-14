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

func TestOpComment(t *testing.T) {
	// success
	assertCode(t, "( comment )")
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
	// success - single true
	assertCode(t, "1 if 123 then", 123)

	// success - single false
	assertCode(t, "0 if 123 then")

	// success - else true
	assertCode(t, "1 if 123 else 456 then", 123)

	// success - else false
	assertCode(t, "0 if 123 else 456 then", 456)
}

func TestOpLoopAndBreak(t *testing.T) {
	// success
	assertCode(t, "loop 123 break done", 123)
}
