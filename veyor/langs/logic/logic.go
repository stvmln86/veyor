// Package logic implements parsing & evaluation functions.
package logic

import (
	"fmt"
	"strings"

	"github.com/stvmln86/veyor/veyor/atoms/atom"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
	"github.com/stvmln86/veyor/veyor/opers"
)

// Evaluate evaluates an Atom against a Queue and Stack.
func Evaluate(a atom.Atom, q *queue.Queue, s *stack.Stack) error {
	switch a := a.(type) {
	case cell.Cell:
		s.Push(a)
		return nil

	case word.Word:
		f, ok := opers.Opers[a]
		if !ok {
			return fmt.Errorf("invalid operator %q", a)
		}

		return f(q, s)

	default:
		return fmt.Errorf("invalid Atom type %T", a)
	}
}

// EvaluateQueue evaluates a Queue against a Stack.
func EvaluateQueue(q *queue.Queue, s *stack.Stack) error {
	for !q.Empty() {
		a, err := q.Dequeue()
		if err != nil {
			return err
		}

		if err := Evaluate(a, q, s); err != nil {
			return err
		}
	}

	return nil
}

// Parse returns an Atom slice from a string.
func Parse(s string) ([]atom.Atom, error) {
	var as []atom.Atom

	for _, s := range strings.Fields(s) {
		a, err := atom.Atomise(s)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}
