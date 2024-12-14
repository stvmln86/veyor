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

	Opers[as[0].(string)] = func(q *Queue, s *Stack) {
		EvaluateSlice(as[1:], s)
	}
}

// OpIf evaluates a conditional if the top Stack integer is true.
func OpIf(q *Queue, s *Stack) {
	var as0, as1 []any
	if q.Contains("else") {
		as1 = q.DequeueTo("else")
		as0 = q.DequeueTo("then")
	} else {
		as1 = q.DequeueTo("then")
	}

	if s.Pop() != 0 {
		EvaluateSlice(as1, s)
	} else {
		EvaluateSlice(as0, s)
	}
}

// OpLoop evaluates a loop until broken.
func OpLoop(q *Queue, s *Stack) {
	as := q.DequeueTo("done")

	for {
		EvaluateSlice(as, s)

		if Break {
			Break = false
			break
		}
	}
}
