package veyor

// OpAdd pushes the sum of the top two Stack integers.
func OpAdd(q *Queue, s *Stack) {
	s.Push(s.Pop() + s.Pop())
}

// OpDivide pushes the quotient of the top two Stack integers.
func OpDivide(q *Queue, s *Stack) {
	s.Push(s.Pop() / s.Pop())
}

// OpDup duplicates the top Stack integer.
func OpDup(q *Queue, s *Stack) {
	s.Push(s.Peek())
}

// OpEq pushes a boolean based on the top Stack integer.
func OpEq(q *Queue, s *Stack) {
	if s.Pop() == s.Pop() {
		s.Push(1)
	} else {
		s.Push(0)
	}
}

// OpLen pushes the number of integers on the Stack.
func OpLen(q *Queue, s *Stack) {
	s.Push(s.Len())
}

// OpModulo pushes the modulo remainder of the top two Stack integers.
func OpModulo(q *Queue, s *Stack) {
	s.Push(s.Pop() % s.Pop())
}

// OpMultiply pushes the product of the top two Stack integers.
func OpMultiply(q *Queue, s *Stack) {
	s.Push(s.Pop() * s.Pop())
}

// OpSubtract pushes the difference of the top two Stack integers.
func OpSubtract(q *Queue, s *Stack) {
	s.Push(s.Pop() - s.Pop())
}

// OpSwap swaps the top two Stack integers.
func OpSwap(q *Queue, s *Stack) {
	s.Push(s.Pop(), s.Pop())
}
