package veyor

import (
	"fmt"
	"strconv"
	"strings"
)

// Oper is a function that operates on a Queue and Stack.
type Oper func(*Queue, *Stack)

// Opers is a map of all defined operator functions.
var Opers map[string]Oper

// init initialises the Opers map.
func init() {
	Opers = map[string]Oper{
		// opers_basic.go
		"+":    OpAdd,
		"/":    OpDivide,
		"%":    OpModulo,
		"*":    OpMultiply,
		"-":    OpSubtract,
		"dup":  OpDup,
		"eq?":  OpEq,
		"len":  OpLen,
		"swap": OpSwap,

		// opers_block.go
		"(":     OpComment,
		"break": OpBreak,
		"def":   OpDef,
		"if":    OpIf,
		"loop":  OpLoop,

		// opers_stdio.go
		"dump":  OpDump,
		"eval":  OpEval,
		"input": OpInput,
		"print": OpPrint,
	}
}

// Evaluate evaluates a Queue against a Stack.
func Evaluate(q *Queue, s *Stack) {
	for !q.Empty() {
		switch a := q.Dequeue().(type) {
		case int:
			s.Push(a)
		case string:
			f, ok := Opers[a]
			if !ok {
				panic(fmt.Sprintf("invalid operator %q", a))
			}

			f(q, s)
		}
	}
}

// EvaluateSlice evaluates an atom slice against a Stack.
func EvaluateSlice(as []any, s *Stack) {
	q := NewQueue(as...)
	Evaluate(q, s)
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
