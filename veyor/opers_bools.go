package veyor

// OpEq pushes a boolean based on the top Stack integer.
func OpEq(q *Queue, s *Stack) {
	if s.Pop() == s.Pop() {
		s.Push(1)
	} else {
		s.Push(0)
	}
}

// OpNot inverts an integer.
func OpNot(q *Queue, s *Stack) {
	if s.Pop() == 0 {
		s.Push(1)
	} else {
		s.Push(0)
	}
}

// OpNoOp does nothing.
func OpNoOp(q *Queue, s *Stack) {
	// nothing
}
