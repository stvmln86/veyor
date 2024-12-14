package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

func TestDef0(t *testing.T) {
	// setup
	q := queue.New(word.Word("abc"), cell.Cell(123), word.Word("end"))
	s := stack.New()

	// success - oper exists
	err := Def0(q, s)
	assert.Contains(t, Opers, word.Word("abc"))
	assert.NoError(t, err)

	// success - oper works
	err = Opers["abc"](nil, s)
	assert.Equal(t, []cell.Cell{123}, s.Cells)
	assert.NoError(t, err)

	// setup
	q = queue.New()

	// error - "def" block missing "end"
	err = Def0(q, nil)
	assert.EqualError(t, err, `"def" block missing "end"`)

	// setup
	q = queue.New(word.Word("end"))

	// error - "def" block missing name/body
	err = Def0(q, nil)
	assert.EqualError(t, err, `"def" block missing name/body`)

	// setup
	q = queue.New(cell.Cell(123), cell.Cell(456), word.Word("end"))

	// error - "def" block name wrong type
	err = Def0(q, nil)
	assert.EqualError(t, err, `"def" block name wrong type`)
}

func TestIf1(t *testing.T) {
	// setup
	q := queue.New(cell.Cell(123), word.Word("then"))
	s := stack.New(1)

	// success - true
	err := If1(q, s)
	assert.Equal(t, []cell.Cell{123}, s.Cells)
	assert.NoError(t, err)

	// setup
	q = queue.New(cell.Cell(123), word.Word("then"))
	s = stack.New(0)

	// success - false
	err = If1(q, s)
	assert.Empty(t, s.Cells)
	assert.NoError(t, err)

	// setup
	q = queue.New()

	// error - "if" block missing "then"
	err = If1(q, nil)
	assert.EqualError(t, err, `"if" block missing "then"`)
}
