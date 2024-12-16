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
//                            part two · helper functions                            //
///////////////////////////////////////////////////////////////////////////////////////

// Try panics on a true boolean with a formatted error message.
func Try(b bool, s string, as ...any) {
	if b {
		panic(fmt.Sprintf(s, as...))
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                         part three · collection functions                         //
///////////////////////////////////////////////////////////////////////////////////////

// Dequeue removes and returns the first atom in a queue.
func Dequeue(as *[]any) any {
	Try(len(*as) == 0, "queue is empty")

	a := (*as)[0]
	(*as) = (*as)[1:]
	return a
}

// DequeueTo removes and returns all atoms up to an atom in a queue.
func DequeueTo(as *[]any, a any) []any {
	i := slices.Index(*as, a)
	Try(i == -1, "queue is missing %v", a)
	as2 := (*as)[:i]
	*as = (*as)[i+1:]
	return as2
}

// Peek returns the last integer in a stack.
func Peek(is *[]int) int {
	Try(len(*is) == 0, "stack is empty")
	return (*is)[len(*is)-1]
}

// Pop removes and returns the last integer in a stack.
func Pop(is *[]int) int {
	Try(len(*is) == 0, "stack is empty")
	i := (*is)[len(*is)-1]
	*is = (*is)[:len(*is)-1]
	return i
}

// Push appends one or more integers to the end of a stack.
func Push(is *[]int, xs ...int) {
	*is = append(*is, xs...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                    part four · parsing and evaluating functions                   //
///////////////////////////////////////////////////////////////////////////////////////

// Evaluate evaluates an queue against a stack.
func Evaluate(as *[]any, is *[]int) {
	for len(*as) > 0 {
		switch a := Dequeue(as).(type) {
		case int:
			Push(is, a)
		case string:
			f, ok := Opers[a]
			Try(!ok, "invalid operator %q", a)
			f(as, is)
		}
	}
}

// EvaluateString evaluates a string against a stack.
func EvaluateString(s string, is *[]int) {
	as := Parse(s)
	Evaluate(&as, is)
}

// Parse returns a queue from a string.
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

// OpAdd ( a b -- c ) adds the last two integers in a stack.
func OpAdd(as *[]any, is *[]int) {
	Push(is, Pop(is)+Pop(is))
}

// OpDivide ( a b -- c ) divides the last two integers in a stack.
func OpDivide(as *[]any, is *[]int) {
	Push(is, Pop(is)/Pop(is))
}

// OpModulo ( a b -- c ) modulos the last two integers in a stack.
func OpModulo(as *[]any, is *[]int) {
	Push(is, Pop(is)%Pop(is))
}

// OpMultiply ( a b -- c ) multiplies the last two integers in a stack.
func OpMultiply(as *[]any, is *[]int) {
	Push(is, Pop(is)*Pop(is))
}

// OpSubtract ( a b -- c ) subtracts the last two integers in a stack.
func OpSubtract(as *[]any, is *[]int) {
	Push(is, Pop(is)-Pop(is))
}

// part five-two · stack operators
///////////////////////////////////

// OpDup ( a -- a a ) duplicates the last integer in a stack.
func OpDup(as *[]any, is *[]int) {
	Push(is, Peek(is))
}

// OpLen ( -- a ) pushes the length of a stack.
func OpLen(as *[]any, is *[]int) {
	Push(is, len(*is))
}

// OpRot ( a b c -- b c a ) rotates the last three integer in a stack.
func OpRot(as *[]any, is *[]int) {
	a, b, c := Pop(is), Pop(is), Pop(is)
	Push(is, b, a, c)
}

// OpSwap ( a b -- b a ) swaps the last two integers in a stack.
func OpSwap(as *[]any, is *[]int) {
	Push(is, Pop(is), Pop(is))
}

// part five-three · block operators
/////////////////////////////////////

// OpAssert ( a -- ) panics if a stack doesn't match the given atoms.
func OpAssert(as *[]any, _ *[]int) {
	is1 := new([]int)
	is2 := new([]int)
	as1 := DequeueTo(as, "=>")
	for _, a := range DequeueTo(as, "end") {
		*is2 = append(*is2, a.(int))
	}

	Evaluate(&as1, is1)
	Try(!slices.Equal(*is1, *is2), "assert: %v => %v\n", as1, *is2)
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
	Try(len(xs) < 2, `"def" missing name/body`)

	_, ok := xs[0].(string)
	Try(!ok, `"def" name is wrong type`)

	Opers[xs[0].(string)] = func(as *[]any, is *[]int) {
		xs := xs[1:]
		Evaluate(&xs, is)
	}
}

// OpIf ( a -- ) evaluates a conditional if the last integer in a stack is true.
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

// part five-four · i/o & eval operators
/////////////////////////////////////////

// OpDump ( -- ) prints a stack to Stdout.
func OpDump(as *[]any, is *[]int) {
	fmt.Fprintf(Stdout, ": %v\n", *is)
}

// OpEval ( ... -- ) evaluates a stack as text up to an EOF zero.
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

// OpInput ( -- ... ) pushes a line from Stdin to a stack.
func OpInput(as *[]any, is *[]int) {
	r := bufio.NewReader(Stdin)
	bs, _ := r.ReadBytes('\n')
	slices.Reverse(bs)

	Push(is, 0)
	for _, b := range bs {
		Push(is, int(b))
	}
}

// OpPrint ( a -- ) prints the last integer in a stack as Unicode.
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
		( a -- b ) ( Negate the last stack integer. )
		dup 2 * · swap -
	end

	( ** Conditional Functions ** )

	def eq?
		( a b -- c ) ( Push 1 if the last two stack integers are equal. )
		- · if 0 else 1 then
	end

	def even?
		( a -- b ) ( Push 1 if the last stack integer is even. )
		2 swap % · zero? if 1 else 0 then
	end

	def neq?
		( a b -- c ) ( Push 1 if the last two integers in a stack are not equal. )
		- · if 1 else 0 then
	end

	def odd?
		( a -- b ) ( Push 1 if the last stack integer is odd. )
		2 swap % · zero? if 0 else 1 then
	end

	def zero?
		( a -- b ) ( Push 1 if the last Stack integer is zero. )
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
		( ... -- ) ( Drop all stack integers. )
		loop · len zero? if break else drop · then · done
	end

	def drop
		( a -- ) ( Drop the last stack integer. )
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
