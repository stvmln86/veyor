package veyor

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

// OpEval evaluates the Stack up to an EOF zero.
func OpEval(q *Queue, st *Stack) {
	var rs []rune
	for !st.Empty() {
		if i := st.Pop(); i == 0 {
			break
		} else {
			rs = append(rs, rune(i))
		}
	}

	if s := strings.TrimSpace(string(rs)); s != "" {
		EvaluateString(s, st)
	}
}

// OpInput pushes a line from Stdin ending with an EOF zero.
func OpInput(q *Queue, s *Stack) {
	r := bufio.NewReader(Stdin)
	bs, _ := r.ReadBytes('\n')
	slices.Reverse(bs)

	s.Push(0)
	for _, b := range bs {
		s.Push(int(b))
	}
}

// OpPrint prints the top Stack integer.
func OpPrint(q *Queue, s *Stack) {
	fmt.Fprintf(Stdout, "%c", s.Pop())
}
