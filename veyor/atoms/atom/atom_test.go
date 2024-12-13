package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomise(t *testing.T) {
	// error - invalid Atom
	a, err := Atomise("")
	assert.Nil(t, a)
	assert.EqualError(t, err, `invalid Atom ""`)
}
