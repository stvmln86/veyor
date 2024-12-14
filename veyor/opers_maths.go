package veyor

// OpAdd pushes the sum of the top two Stack integers.
func OpAdd(q *Queue, s *Stack) {
	s.Push(s.Pop() + s.Pop())
}

// OpDivide pushes the quotient of the top two Stack integers.
func OpDivide(q *Queue, s *Stack) {
	s.Push(s.Pop() / s.Pop())
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
