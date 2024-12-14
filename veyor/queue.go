package veyor

import "slices"

// Queue is a first-in-first-out queue of atoms.
type Queue struct {
	Atoms []any
}

// NewQueue returns a new Queue from zero or more atoms.
func NewQueue(as ...any) *Queue {
	return &Queue{as}
}

// Contains returns true if the Queue contains an atom.
func (q *Queue) Contains(a any) bool {
	return slices.Contains(q.Atoms, a)
}

// Dequeue removes and returns the first atom in the Queue.
func (q *Queue) Dequeue() any {
	if len(q.Atoms) == 0 {
		panic("Queue is empty")
	}

	a := q.Atoms[0]
	q.Atoms = q.Atoms[1:]
	return a
}

// DequeueTo removes and returns all atoms up to an atom in the Queue.
func (q *Queue) DequeueTo(a any) []any {
	i := slices.Index(q.Atoms, a)
	if i == -1 {
		panic("Queue is insufficient")
	}

	as := q.Atoms[:i]
	q.Atoms = q.Atoms[i+1:]
	return as
}

// Empty returns true if the Queue has no atoms.
func (q *Queue) Empty() bool {
	return len(q.Atoms) == 0
}

// Enqueue appends one or more atoms to the end of the Queue.
func (q *Queue) Enqueue(as ...any) {
	q.Atoms = append(q.Atoms, as...)
}

// Len returns the number of atoms in the Queue.
func (q *Queue) Len() int {
	return len(q.Atoms)
}
