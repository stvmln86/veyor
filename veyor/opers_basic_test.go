package veyor

import "testing"

func TestOpAdd(t *testing.T) {
	// success
	assertCode(t, "1 2 +", 3)
}

func TestOpDivide(t *testing.T) {
	// success
	assertCode(t, "3 6 /", 2)
}

func TestOpDup(t *testing.T) {
	// success
	assertCode(t, "123 dup", 123, 123)
}

func TestOpEq(t *testing.T) {
	// success
	assertCode(t, "1 2 eq?", 0)
	assertCode(t, "1 1 eq?", 1)
}

func TestOpLen(t *testing.T) {
	// success
	assertCode(t, "123 len", 123, 1)
}

func TestOpModulo(t *testing.T) {
	// success
	assertCode(t, "3 5 %", 2)
}

func TestOpMultiply(t *testing.T) {
	// success
	assertCode(t, "2 3 *", 6)
}

func TestOpSubtract(t *testing.T) {
	// success
	assertCode(t, "3 5 -", 2)
}

func TestOpSwap(t *testing.T) {
	// success
	assertCode(t, "123 456 swap", 456, 123)
}
