// Package stack implements the Stack type and methods.
package stack

import (
	"fmt"
	"strconv"
	"strings"
)

// Stack is a last-in-first-out stack of integers.
type Stack struct {
	Items []int
}

// New returns a new Stack from zero or more integers.
func New(is ...int) *Stack {
	return &Stack{is}
}

// Empty returns true if the Stack is empty.
func (s *Stack) Empty() bool {
	return len(s.Items) == 0
}

// Len returns the number of integers in the Stack.
func (s *Stack) Len() int {
	return len(s.Items)
}

// Pop removes and returns the top integer on the Stack.
func (s *Stack) Pop() (int, error) {
	if len(s.Items) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}

	i := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return i, nil
}

// Push appends an integer to the top of the Stack.
func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

// String returns the Stack as a string.
func (s *Stack) String() string {
	var ss []string
	for _, i := range s.Items {
		ss = append(ss, strconv.Itoa(i))
	}

	return strings.Join(ss, " ")
}
