// Package opers implements the Oper func type and definitions.
package opers

import (
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
	}
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
