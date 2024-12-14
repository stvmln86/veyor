package veyor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	// success
	q := NewQueue(123)
	assert.Equal(t, []any{123}, q.Atoms)
}

func TestQueueContains(t *testing.T) {
	// success - true
	ok := NewQueue(123).Contains(123)
	assert.True(t, ok)

	// success - false
	ok = NewQueue().Contains(123)
	assert.False(t, ok)
}

func TestQueueDequeue(t *testing.T) {
	// setup
	q := NewQueue(123)

	// success
	a := q.Dequeue()
	assert.Equal(t, 123, a)
	assert.Empty(t, q.Atoms)

	// error - Queue is empty
	defer func() { assert.Equal(t, "Queue is empty", recover()) }()
	q.Dequeue()
}

func TestQueueDequeueTo(t *testing.T) {
	// setup
	q := NewQueue(123, "end")

	// success
	as := q.DequeueTo("end")
	assert.Equal(t, []any{123}, as)
	assert.Empty(t, q.Atoms)

	// error - Queue is insufficient
	defer func() { assert.Equal(t, "Queue is insufficient", recover()) }()
	q.DequeueTo("end")
}

func TestQueueEmpty(t *testing.T) {
	// success - true
	ok := NewQueue().Empty()
	assert.True(t, ok)

	// success - false
	ok = NewQueue(123).Empty()
	assert.False(t, ok)
}

func TestQueueEnqueue(t *testing.T) {
	// setup
	q := NewQueue()

	// success
	q.Enqueue(123, 456)
	assert.Equal(t, []any{123, 456}, q.Atoms)
}

func TestQueueLen(t *testing.T) {
	// success
	i := NewQueue(123).Len()
	assert.Equal(t, 1, i)
}
