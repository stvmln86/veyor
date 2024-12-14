// Package logic implements program logic and operator functions.
package logic

import (
	"fmt"
	"strings"

	"github.com/stvmln86/veyor/veyor/atoms/atom"
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/atoms/word"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

// Oper is a function that operates on a Queue and Stack.
type Oper func(*queue.Queue, *stack.Stack) error

// Opers is a map of all existing Oper functions.
var Opers map[word.Word]Oper

// init initialises the Opers map.
func init() {
	Opers = map[word.Word]Oper{
		// conds.go
		"if":  If1,
		"def": Def0,

		// maths.go
		"+": Wrap(2, Add2),
		"/": Wrap(2, Div2),
		"%": Wrap(2, Mod2),
		"*": Wrap(2, Mul2),
		"-": Wrap(2, Sub2),

		// stack.go
		"dupe": Wrap(1, Dupe1),
		"drop": Wrap(1, Drop1),
		"roll": Wrap(3, Roll3),
		"swap": Wrap(2, Swap2),

		// stdio.go
		"input": Input0,
		"print": PrintN,
	}
}

// Evaluate evaluates an Atom against a Queue and Stack.
func Evaluate(a atom.Atom, q *queue.Queue, s *stack.Stack) error {
	switch a := a.(type) {
	case cell.Cell:
		s.Push(a)
		return nil

	case word.Word:
		f, ok := Opers[a]
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

// Wrap returns a Cell slice function as an Oper.
func Wrap(n int, f func([]cell.Cell) []cell.Cell) Oper {
	return func(_ *queue.Queue, s *stack.Stack) error {
		cs, err := s.PopN(n)
		if err != nil {
			return err
		}

		s.PushAll(f(cs))
		return nil
	}
}
