package veyor

import "testing"

func TestOpEq(t *testing.T) {
	// success
	assertCode(t, "1 2 eq?", 0)
	assertCode(t, "1 1 eq?", 1)
}

func TestOpNot(t *testing.T) {
	// success
	assertCode(t, "1 not", 0)
	assertCode(t, "0 not", 1)
}

func TestOpNoOp(t *testing.T) {
	// success
	assertCode(t, "·")
}
