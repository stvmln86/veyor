package veyor

import (
	"strconv"
	"strings"
)

// Evaluate evaluates a Queue against a Stack.
func Evaluate(q *Queue, s *Stack) {
	for !q.Empty() {
		switch a := q.Dequeue().(type) {
		case int:
			s.Push(a)
		case string:
			f, ok := Opers[a]
			if !ok {
				panic("invalid operator")
			}

			f(q, s)
		}
	}
}

// EvaluateString evaluates a string against a Stack.
func EvaluateString(s string, st *Stack) {
	q := NewQueue(Parse(s)...)
	Evaluate(q, st)
}

// Parse returns an atom slice from a string.
func Parse(s string) []any {
	var as []any

	for _, s := range strings.Fields(s) {
		i, err := strconv.Atoi(s)
		if err == nil {
			as = append(as, i)
		} else {
			as = append(as, s)
		}
	}

	return as
}
