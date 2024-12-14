package veyor

import "testing"

func TestOpDrop(t *testing.T) {
	// success
	assertCode(t, "123 drop")
}

func TestOpDup(t *testing.T) {
	// success
	assertCode(t, "123 dup", 123, 123)
}

func TestOpLen(t *testing.T) {
	// success
	assertCode(t, "123 len", 123, 1)
}

func TestOpSwap(t *testing.T) {
	// success
	assertCode(t, "123 456 swap", 456, 123)
}
