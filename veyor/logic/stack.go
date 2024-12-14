package logic

import (
	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

// Dupe1 duplicates the top Cell on the Stack.
func Dupe1(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0], cs[0]}
}

// Drop1 deletes the top Cell on the Stack.
func Drop1(cs []cell.Cell) []cell.Cell {
	return nil
}

// Len0 returns the number of Cells on the Stack.
func Len0(_ *queue.Queue, s *stack.Stack) error {
	s.Push(cell.Cell(s.Len()))
	return nil
}

// Roll3 rotates the top three Cells on the Stack.
func Roll3(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[1], cs[0], cs[2]}
}

// Swap2 swaps the top two Cells on the Stack.
func Swap2(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0], cs[1]}
}
