package opers

import "github.com/stvmln86/veyor/veyor/atoms/cell"

// Add2 pushes the sum of the top two Cells on a Stack.
func Add2(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0] + cs[1]}
}

// Div2 pushes the quotient of the top two Cells on a Stack.
func Div2(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0] / cs[1]}
}

// Mod2 pushes the modulo remainder of the top two Cells on a Stack.
func Mod2(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0] % cs[1]}
}

// Mul2 pushes the product of the top two Cells on a Stack.
func Mul2(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0] * cs[1]}
}

// Sub2 pushes the difference of the top two Cells on a Stack.
func Sub2(cs []cell.Cell) []cell.Cell {
	return []cell.Cell{cs[0] - cs[1]}
}
