// Package queue implements the Queue type.
package queue

import (
	"fmt"
	"slices"

	"github.com/stvmln86/veyor/veyor/atoms/atom"
)

// Queue is a first-in-first-out queue of Atoms.
type Queue struct {
	Atoms []atom.Atom
}

// New returns a new Queue from zero or more Atoms.
func New(as ...atom.Atom) *Queue {
	return &Queue{as}
}

// Dequeue removes and returns the first Atom in the Queue.
func (q *Queue) Dequeue() (atom.Atom, error) {
	if len(q.Atoms) == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	a := q.Atoms[0]
	q.Atoms = q.Atoms[1:]
	return a, nil
}

// DequeueTo removes and returns all Atoms up to an Atom in the Queue.
func (q *Queue) DequeueTo(a atom.Atom) ([]atom.Atom, error) {
	i := slices.Index(q.Atoms, a)
	if i == -1 {
		return nil, fmt.Errorf("Queue is insufficient")
	}

	as := q.Atoms[:i]
	q.Atoms = q.Atoms[i+1:]
	return as, nil
}

// Empty returns true if the Queue has no Atoms.
func (q *Queue) Empty() bool {
	return len(q.Atoms) == 0
}

// Enqueue appends an Atom to the end of the Queue.
func (q *Queue) Enqueue(a atom.Atom) {
	q.Atoms = append(q.Atoms, a)
}

// EnqueueAll appends an Atom slice to the end of the Queue.
func (q *Queue) EnqueueAll(as []atom.Atom) {
	q.Atoms = append(q.Atoms, as...)
}

// Len returns the number of Atoms in the Queue.
func (q *Queue) Len() int {
	return len(q.Atoms)
}
