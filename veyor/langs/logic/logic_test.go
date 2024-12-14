package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/atom"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

func TestEvaluate(t *testing.T) {
	// setup
	s := stack.New(1)

	// success - Cell
	err := Evaluate(cell.Cell(2), nil, s)
	assert.Equal(t, []cell.Cell{1, 2}, s.Cells)
	assert.NoError(t, err)

	// success - Word
	err = Evaluate(word.Word("+"), nil, s)
	assert.Equal(t, []cell.Cell{3}, s.Cells)
	assert.NoError(t, err)

	// error - invalid Atom type
	err = Evaluate(nil, nil, nil)
	assert.EqualError(t, err, `invalid Atom type <nil>`)
}

func TestEvaluateQueue(t *testing.T) {
	// setup
	q := queue.New(cell.Cell(2), word.Word("+"))
	s := stack.New(1)

	// success
	err := EvaluateQueue(q, s)
	assert.Equal(t, []cell.Cell{3}, s.Cells)
	assert.NoError(t, err)
}

func TestParse(t *testing.T) {
	// success
	as, err := Parse("123 abc")
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, as)
	assert.NoError(t, err)
}
