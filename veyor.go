///////////////////////////////////////////////////////////////////////////////////////
//                  veyor · a minimal stack language in Go · v0.0.0                  //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"flag"
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
type Oper func(*[]any, *[]int)

// Opers is a map of all defined Opers.
var Opers map[string]Oper

///////////////////////////////////////////////////////////////////////////////////////
//                          part two · collection functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// Dequeue removes and returns the first atom in an atom slice.
func Dequeue(as *[]any) any {
	if len(*as) == 0 {
		panic("queue is empty")
	}

	a := (*as)[0]
	(*as) = (*as)[1:]
	return a
}

// DequeueTo removes and returns all atoms up to an atom in an atom slice.
func DequeueTo(as *[]any, a any) []any {
	i := slices.Index(*as, a)
	if i == -1 {
		panic(fmt.Sprintf("queue is missing %v", a))
	}

	as2 := (*as)[:i]
	*as = (*as)[i+1:]
	return as2
}

// Peek returns the top integer on an integer slice.
func Peek(is *[]int) int {
	if len(*is) == 0 {
		panic("stack is empty")
	}

	return (*is)[len(*is)-1]
}

// Pop removes and returns the top integer on an integer slice.
func Pop(is *[]int) int {
	if len(*is) == 0 {
		panic("stack is empty")
	}

	i := (*is)[len(*is)-1]
	*is = (*is)[:len(*is)-1]
	return i
}

// Push appends one or more integers to the top of an integer slice.
func Push(is *[]int, xs ...int) {
	*is = append(*is, xs...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                   part three · parsing and evaluating functions                   //
///////////////////////////////////////////////////////////////////////////////////////

// Evaluate evaluates a Queue against a Stack.
func Evaluate(as *[]any, is *[]int) {
	for len(*as) > 0 {
		switch a := Dequeue(as).(type) {
		case int:
			Push(is, a)
		case string:
			f, ok := Opers[a]
			if !ok {
				panic(fmt.Sprintf("invalid operator %q", a))
			}

			f(as, is)
		}
	}
}

// EvaluateString evaluates a string against a Stack.
func EvaluateString(s string, is *[]int) {
	as := Parse(s)
	Evaluate(&as, is)
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
func OpAdd(as *[]any, is *[]int) { Push(is, Pop(is)+Pop(is)) }

// OpDivide ( a b -- c ) divides the top two Stack integers.
func OpDivide(as *[]any, is *[]int) { Push(is, Pop(is)/Pop(is)) }

// OpModulo ( a b -- c ) modulos the top two Stack integers.
func OpModulo(as *[]any, is *[]int) { Push(is, Pop(is)%Pop(is)) }

// OpMultiply ( a b -- c ) multiplies the top two Stack integers.
func OpMultiply(as *[]any, is *[]int) { Push(is, Pop(is)*Pop(is)) }

// OpSubtract ( a b -- c ) subtracts the top two Stack integers.
func OpSubtract(as *[]any, is *[]int) { Push(is, Pop(is)-Pop(is)) }

// part five-two · stack operators
///////////////////////////////////

// OpDup ( a -- a a ) duplicates the top Stack integer.
func OpDup(as *[]any, is *[]int) {
	Push(is, Peek(is))
}

// OpLen ( -- a ) pushes the Stack length.
func OpLen(as *[]any, is *[]int) {
	Push(is, len(*is))
}

// OpRot ( a b c -- b c a ) rotates the top three Stack integers.
func OpRot(as *[]any, is *[]int) {
	a, b, c := Pop(is), Pop(is), Pop(is)
	Push(is, b, a, c)
}

// OpSwap ( a b -- b a ) swaps the top two Stack integers.
func OpSwap(as *[]any, is *[]int) {
	Push(is, Pop(is), Pop(is))
}

// part five-three · logical operators
///////////////////////////////////////

// OpEq ( a b -- c ) pushes 1 if the top two Stack integers are equal.
func OpEq(as *[]any, is *[]int) {
	if Pop(is) == Pop(is) {
		Push(is, 1)
	} else {
		Push(is, 0)
	}
}

// part five-four · block operators
/////////////////////////////////////

// OpAssert ( a -- ) panics if the Stack doesn't match the given atoms.
func OpAssert(as *[]any, is *[]int) {
	xs := DequeueTo(as, "end")
	slices.Reverse(xs)

	if len(xs) == 0 && len(*is) != 0 {
		panic("assert error · stack should be empty")
	}

	for _, a := range xs {
		if Pop(is) != a.(int) {
			panic(fmt.Sprintf("assert error · stack should be [%v]", xs))
		}
	}
}

// OpBreak ( -- ) breaks the current loop.
func OpBreak(as *[]any, is *[]int) {
	Break = true
}

// OpComment ( -- ) defines a comment.
func OpComment(as *[]any, is *[]int) {
	DequeueTo(as, ")")
}

// OpDef ( -- ) defines a custom operator.
func OpDef(as *[]any, is *[]int) {
	xs := DequeueTo(as, "end")

	if len(xs) < 2 {
		panic(`"def" missing name/body`)
	}

	if _, ok := xs[0].(string); !ok {
		panic(`"def" name is wrong type`)
	}

	Opers[xs[0].(string)] = func(as *[]any, is *[]int) {
		xs := xs[1:]
		Evaluate(&xs, is)
	}
}

// OpIf ( a -- ) evaluates a conditional if the top Stack integer is true.
func OpIf(as *[]any, is *[]int) {
	var as1, as2 []any
	i := slices.Index(*as, "else")

	if i != -1 && i < slices.Index(*as, "then") {
		as2 = DequeueTo(as, "else")
		as1 = DequeueTo(as, "then")
	} else {
		as2 = DequeueTo(as, "then")
	}

	if Pop(is) != 0 {
		Evaluate(&as2, is)
	} else {
		Evaluate(&as1, is)
	}
}

// OpLoop ( -- ) evaluates a loop until broken.
func OpLoop(as *[]any, is *[]int) {
	xs := DequeueTo(as, "done")

	for {
		var xs2 = make([]any, len(xs))
		copy(xs2, xs)
		Evaluate(&xs2, is)
		if Break {
			Break = false
			break
		}
	}
}

// part five-five · i/o & eval operators
/////////////////////////////////////////

// OpDump ( -- ) prints the Stack.
func OpDump(as *[]any, is *[]int) {
	var ss []string
	for _, i := range *is {
		ss = append(ss, strconv.Itoa(i))
	}

	fmt.Fprintf(Stdout, "[ %s ]\n", strings.Join(ss, " "))
}

// OpEval ( ... -- ) evaluates the Stack up to an EOF zero.
func OpEval(as *[]any, is *[]int) {
	var rs []rune
	for len(*is) > 0 {
		if i := Pop(is); i == 0 {
			break
		} else {
			rs = append(rs, rune(i))
		}
	}

	if s := strings.TrimSpace(string(rs)); s != "" {
		EvaluateString(s, is)
	}
}

// OpInput ( -- ... ) pushes a line from Stdin ending with an EOF zero.
func OpInput(as *[]any, is *[]int) {
	r := bufio.NewReader(Stdin)
	bs, _ := r.ReadBytes('\n')
	slices.Reverse(bs)

	Push(is, 0)
	for _, b := range bs {
		Push(is, int(b))
	}
}

// OpPrint ( a -- ) prints the top Stack integer.
func OpPrint(as *[]any, is *[]int) {
	fmt.Fprintf(Stdout, "%c", Pop(is))
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
		"rot":  OpRot,
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
	f := flag.NewFlagSet("veyor", flag.ExitOnError)
	c := f.String("c", "", "execute string and exit")
	f.Parse(os.Args[1:])

	switch {
	case *c != "":
		EvaluateString(Stlib+*c+"\ndump", new([]int))

	default:
		EvaluateString(Stlib+"\nrepl", new([]int))
	}
}
