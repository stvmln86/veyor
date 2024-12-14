package veyor

// Stack is a last-in-first-out stack of integers.
type Stack struct {
	Items []int
}

// NewStack returns a new Stack from zero or more integers.
func NewStack(is ...int) *Stack {
	return &Stack{is}
}

// Empty returns true if the Stack has no integers.
func (s *Stack) Empty() bool {
	return len(s.Items) == 0
}

// Len returns the number of integers on the Stack.
func (s *Stack) Len() int {
	return len(s.Items)
}

// Peek returns the top integer on the Stack.
func (s *Stack) Peek() int {
	if len(s.Items) == 0 {
		panic("Stack is empty")
	}

	return s.Items[len(s.Items)-1]
}

// Pop removes and returns the top integer on the Stack.
func (s *Stack) Pop() int {
	if len(s.Items) == 0 {
		panic("Stack is empty")
	}

	i := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return i
}

// Push appends one or more integers to the top of the Stack.
func (s *Stack) Push(is ...int) {
	s.Items = append(s.Items, is...)
}
