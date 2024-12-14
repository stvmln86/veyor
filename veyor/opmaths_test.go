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
