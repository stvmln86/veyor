package opers

import "github.com/stvmln86/veyor/veyor/atoms/cell"

// Dupe1 duplicates the top Cell on the Stack.
func Dupe1(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0], cs[0]}
}

// Drop1 deletes the top Cell on the Stack.
func Drop1(cs []cell.Cell) []cell.Cell {
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
