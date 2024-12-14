package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/veyor/veyor/atoms/atom"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
)

func TestNew(t *testing.T) {
	// success
	q := New(cell.Cell(123), word.Word("abc"))
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, q.Atoms)
}

func TestDequeue(t *testing.T) {
	// setup
	q := New(cell.Cell(123))

	// success
	a, err := q.Dequeue()
	assert.Equal(t, cell.Cell(123), a)
	assert.Empty(t, q.Atoms)
	assert.NoError(t, err)

	// error - Queue is empty
	a, err = q.Dequeue()
	assert.Zero(t, a)
	assert.EqualError(t, err, "Queue is empty")
}

func TestDequeueTo(t *testing.T) {
	// setup
	q := New(word.Word("abc"), word.Word("end"))

	// success
	as, err := q.DequeueTo(word.Word("end"))
	assert.Equal(t, []atom.Atom{word.Word("abc")}, as)
	assert.Empty(t, q.Atoms)
	assert.NoError(t, err)

	// error - Queue is insufficient
	as, err = q.DequeueTo(word.Word("end"))
	assert.Empty(t, as)
	assert.EqualError(t, err, "Queue is insufficient")
}

func TestEmpty(t *testing.T) {
	// success - true
	ok := New().Empty()
	assert.True(t, ok)

	// success - true
	ok = New(cell.Cell(123), word.Word("abc")).Empty()
	assert.False(t, ok)
}

func TestLen(t *testing.T) {
	// success
	n := New(cell.Cell(123), word.Word("abc")).Len()
	assert.Equal(t, 2, n)
}

func TestEnqueue(t *testing.T) {
	// setup
	q := New()

	// success
	q.Enqueue(cell.Cell(123))
	assert.Equal(t, []atom.Atom{cell.Cell(123)}, q.Atoms)
}

func TestPushAll(t *testing.T) {
	// setup
	q := New()

	// success
	q.EnqueueAll([]atom.Atom{cell.Cell(123), word.Word("abc")})
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, q.Atoms)
}
