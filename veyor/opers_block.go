package veyor

// OpBreak breaks the current loop.
func OpBreak(q *Queue, s *Stack) {
	Break = true
}

// OpComment defines a comment.
func OpComment(q *Queue, s *Stack) {
	q.DequeueTo(")")
}

// OpDef defines a custom operator function.
func OpDef(q *Queue, s *Stack) {
	as := q.DequeueTo("end")

	if len(as) < 2 {
		panic(`"def" missing name/body`)
	}

	if _, ok := as[0].(string); !ok {
		panic(`"def" name is wrong type`)
	}

	q2 := NewQueue(as...)
	Opers[as[0].(string)] = func(q *Queue, s *Stack) {
		Evaluate(q2, s)
	}
}

// OpIf evaluates a conditional if the top Stack integer is not zero.
func OpIf(q *Queue, s *Stack) {
	as := q.DequeueTo("then")
	q2 := NewQueue(as...)

	if s.Pop() != 0 {
		Evaluate(q2, s)
	}
}

// OpLoop evaluates a loop until broken.
func OpLoop(q *Queue, s *Stack) {
	as := q.DequeueTo("done")
	q2 := NewQueue(as...)

	for {
		Evaluate(q2, s)

		if Break {
			Break = false
			break
		}
	}
}
