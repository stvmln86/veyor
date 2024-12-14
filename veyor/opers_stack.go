package veyor

// OpDrop removes the top Stack integer.
func OpDrop(q *Queue, s *Stack) {
	s.Pop()
}

// OpDup duplicates the top Stack integer.
func OpDup(q *Queue, s *Stack) {
	s.Push(s.Peek())
}

// OpLen pushes the number of integers on the Stack.
func OpLen(q *Queue, s *Stack) {
	s.Push(s.Len())
}

// OpSwap swaps the top two Stack integers.
func OpSwap(q *Queue, s *Stack) {
	s.Push(s.Pop(), s.Pop())
}
