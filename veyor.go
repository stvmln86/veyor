///////////////////////////////////////////////////////////////////////////////////////
//                  veyor · a minimal stack language in Go · v0.0.0                  //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                         //
///////////////////////////////////////////////////////////////////////////////////////

// Break indicates that the current loop should exit.
var Break bool = false

// Stdin is the default input stream.
var Stdin io.Reader = os.Stdin

// Stdout is the default output stream.
var Stdout io.Writer = os.Stdout

// Oper is a callable program function.
type Oper func(*Queue, *Stack)

// Opers is a map of all defined Opers.
var Opers map[string]Oper

///////////////////////////////////////////////////////////////////////////////////////
//                             part two · the queue type                             //
///////////////////////////////////////////////////////////////////////////////////////

// Queue is a first-in-first-out queue of atoms.
type Queue struct {
	Atoms []any
}

// NewQueue returns a new Queue from zero or more atoms.
func NewQueue(as ...any) *Queue {
	return &Queue{as}
}

// Dequeue removes and returns the first atom in the Queue.
func (q *Queue) Dequeue() any {
	if len(q.Atoms) == 0 {
		panic("Queue is empty")
	}

	a := q.Atoms[0]
	q.Atoms = q.Atoms[1:]
	return a
}

// DequeueTo removes and returns all atoms up to an atom in the Queue.
func (q *Queue) DequeueTo(a any) []any {
	i := slices.Index(q.Atoms, a)
	if i == -1 {
		panic(fmt.Sprintf("Queue is missing %v", a))
	}

	as := q.Atoms[:i]
	q.Atoms = q.Atoms[i+1:]
	return as
}

// Empty returns true if the Queue has no atoms.
func (q *Queue) Empty() bool {
	return len(q.Atoms) == 0
}

// Enqueue appends one or more atoms to the end of the Queue.
func (q *Queue) Enqueue(as ...any) {
	q.Atoms = append(q.Atoms, as...)
}

// Index returns the index of an atom in the Queue.
func (q *Queue) Index(a any) int {
	return slices.Index(q.Atoms, a)
}

// Len returns the number of atoms in the Queue.
func (q *Queue) Len() int {
	return len(q.Atoms)
}

// String returns the Queue as a string.
func (q *Queue) String() string {
	var ss []string
	for _, a := range q.Atoms {
		ss = append(ss, fmt.Sprintf("%v", a))
	}

	return strings.Join(ss, " ")
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part three · the stack type                            //
///////////////////////////////////////////////////////////////////////////////////////

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

// String returns the Stack as a string.
func (s *Stack) String() string {
	var ss []string
	for _, a := range s.Items {
		ss = append(ss, strconv.Itoa(a))
	}

	return strings.Join(ss, " ")
}

///////////////////////////////////////////////////////////////////////////////////////
//                    part four · parsing and evaluating functions                   //
///////////////////////////////////////////////////////////////////////////////////////

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
		if i, err := strconv.Atoi(s); err == nil {
			as = append(as, i)
		} else {
			as = append(as, s)
		}
	}

	return as
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part five · operator functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// part five-one · math operators
//////////////////////////////////

// OpAdd ( a b -- c ) adds the top two Stack integers.
func OpAdd(q *Queue, s *Stack) {
	s.Push(s.Pop() + s.Pop())
}

// OpDivide ( a b -- c ) divides the top two Stack integers.
func OpDivide(q *Queue, s *Stack) {
	s.Push(s.Pop() / s.Pop())
}

// OpModulo ( a b -- c ) modulos the top two Stack integers.
func OpModulo(q *Queue, s *Stack) {
	s.Push(s.Pop() % s.Pop())
}

// OpMultiply ( a b -- c ) multiplies the top two Stack integers.
func OpMultiply(q *Queue, s *Stack) {
	s.Push(s.Pop() * s.Pop())
}

// OpSubtract ( a b -- c ) subtracts the top two Stack integers.
func OpSubtract(q *Queue, s *Stack) {
	s.Push(s.Pop() - s.Pop())
}

// part five-two · stack operators
///////////////////////////////////

// OpDup ( a -- a a ) duplicates the top Stack integer.
func OpDup(q *Queue, s *Stack) {
	s.Push(s.Peek())
}

// OpLen ( -- a ) pushes the Stack length.
func OpLen(q *Queue, s *Stack) {
	s.Push(s.Len())
}

// OpSwap ( a b -- b a ) swaps the top two Stack integers.
func OpSwap(q *Queue, s *Stack) {
	s.Push(s.Pop(), s.Pop())
}

// part five-three · logical operators
///////////////////////////////////////

// OpEq ( a b -- c ) pushes 1 if the top two Stack integers are equal.
func OpEq(q *Queue, s *Stack) {
	if s.Pop() == s.Pop() {
		s.Push(1)
	} else {
		s.Push(0)
	}
}

// part five-four · block operators
/////////////////////////////////////

// OpAssert ( a -- ) panics if the Stack doesn't match the given atoms.
func OpAssert(q *Queue, s *Stack) {
	as := q.DequeueTo("end")
	slices.Reverse(as)

	for _, a := range as {
		if s.Pop() != a.(int) {
			q := NewQueue(as...)
			panic(fmt.Sprintf("assert error · stack should be [%s]", q))
		}
	}
}

// OpBreak ( -- ) breaks the current loop.
func OpBreak(q *Queue, s *Stack) {
	Break = true
}

// OpComment ( -- ) defines a comment.
func OpComment(q *Queue, s *Stack) {
	q.DequeueTo(")")
}

// OpDef ( -- ) defines a custom operator.
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

// OpIf ( a -- ) evaluates a conditional if the top Stack integer is true.
func OpIf(q *Queue, s *Stack) {
	var as0, as1 []any
	i := q.Index("else")

	if i != -1 && i < q.Index("then") {
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

// OpLoop ( -- ) evaluates a loop until broken.
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

// part five-five · i/o & eval operators
/////////////////////////////////////////

// OpDump ( -- ) prints the Stack.
func OpDump(q *Queue, s *Stack) {
	fmt.Fprintf(Stdout, "[ %s ]\n", s.String())
}

// OpEval ( ... -- ) evaluates the Stack up to an EOF zero.
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

// OpInput ( -- ... ) pushes a line from Stdin ending with an EOF zero.
func OpInput(q *Queue, s *Stack) {
	r := bufio.NewReader(Stdin)
	bs, _ := r.ReadBytes('\n')
	slices.Reverse(bs)

	s.Push(0)
	for _, b := range bs {
		s.Push(int(b))
	}
}

// OpPrint ( a -- ) prints the top Stack integer.
func OpPrint(q *Queue, s *Stack) {
	fmt.Fprintf(Stdout, "%c", s.Pop())
}

///////////////////////////////////////////////////////////////////////////////////////
//                          part six · the standard library                          //
///////////////////////////////////////////////////////////////////////////////////////

// Stlib is Veyor's standard library.
const Stlib = `
	( ** Boolean Functions ** )

	def not
		( a -- b ) ( Negate the top Stack integer. )
		dup 2 * · swap -
	end

	( ** Conditional Functions ** )

	def even?
		( a -- b ) ( Push 1 if the top Stack integer is even. )
		2 swap % · zero? if 1 else 0 then
	end

	def neq?
		( a b -- c ) ( Push 1 if the top two Stack integers are not equal. )
		eq? · if 0 else 1 then
	end

	def odd?
		( a -- b ) ( Push 1 if the top Stack integer is odd. )
		2 swap % · zero? if 0 else 1 then
	end

	def zero?
		( a -- b ) ( Push 1 if the top Stack integer is zero. )
		0 eq? · if 1 else 0 then
	end

	( ** Interactive Functions ** )

	def repl
		( -- ) ( Launch a read-evaluate-print loop. )
		loop · 0 32 62 print0 · input eval · len if dump then · done
	end

	( ** Miscellaneous Functions ** )

	def ·
		( -- ) ( Do nothing. )
	end

	( ** Stack Functions ** )

	def clear
		( ... -- ) ( Drop all Stack integers. )
		loop · len zero? if break else drop · then · done
	end

	def drop
		( a -- ) ( Drop the top Stack integer. )
		if then
	end

	( ** Standard I/O Functions ** )

	def print0
		( a... -- ) ( Print until an EOF zero. )
		loop · dup zero? if drop break else print then · done
	end
`

///////////////////////////////////////////////////////////////////////////////////////
//                           part seven · the main runtime                           //
///////////////////////////////////////////////////////////////////////////////////////

// init initialises the main Veyor program.
func init() {
	Opers = map[string]Oper{
		// math operators
		"+": OpAdd,
		"/": OpDivide,
		"%": OpModulo,
		"*": OpMultiply,
		"-": OpSubtract,

		// stack operators
		"dup":  OpDup,
		"len":  OpLen,
		"swap": OpSwap,

		// logical operators
		"eq?": OpEq,

		// block operators
		"assert": OpAssert,
		"break":  OpBreak,
		"(":      OpComment,
		"def":    OpDef,
		"if":     OpIf,
		"loop":   OpLoop,

		// i/o & eval operators
		"dump":  OpDump,
		"eval":  OpEval,
		"input": OpInput,
		"print": OpPrint,
	}
}

// main runs the main Veyor program.
func main() {
	s := NewStack()
	EvaluateString(Stlib+"repl", s)
}
