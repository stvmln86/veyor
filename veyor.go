///////////////////////////////////////////////////////////////////////////////////////
//                  veyor · a minimal stack language in Go · v0.0.0                  //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                              part 1 · stack functions                             //
///////////////////////////////////////////////////////////////////////////////////////

// Peek returns the last integer in a stack slice.
func Peek(is *[]int) int {
	if len(*is) == 0 {
		panic("Stack is empty")
	}

	return (*is)[len(*is)-1]
}

// Pop removes and returns the last integer in a stack slice.
func Pop(is *[]int) int {
	if len(*is) == 0 {
		panic("Stack is empty")
	}

	i := (*is)[len(*is)-1]
	*is = (*is)[:len(*is)-1]
	return i
}

// Push appends one or more integers to the end of a stack slice.
func Push(is *[]int, xs ...int) {
	*is = append(*is, xs...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                              part 2 · queue functions                             //
///////////////////////////////////////////////////////////////////////////////////////

// Dequeue removes and returns the first atom in a queue slice.
func Dequeue(as *[]any) any {
	if len(*as) == 0 {
		panic("Queue is empty")
	}

	a := (*as)[0]
	*as = (*as)[1:]
	return a
}

// DequeueTo removes and returns all atoms up to an atom in a queue slice.
func DequeueTo(as *[]any, a any) []any {
	i := slices.Index(*as, a)
	if i == -1 {
		panic("Queue is insufficient")
	}

	xs := (*as)[:i]
	*as = (*as)[i+1:]
	return xs
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part 3 · parsing & evaluating                           //
///////////////////////////////////////////////////////////////////////////////////////

// Evaluate evaluates an atom against a Queue and Stack.
func Evaluate(a any, as *[]any, is *[]int) {
	switch a := a.(type) {
	case int:
		*is = append(*is, a)
	case string:
		f, ok := Opers[a]
		if !ok {
			panic(fmt.Sprintf("invalid operator %q", a))
		}

		f(as, is)
	}
}

// EvaluateQueue evaluates a Queue against a stack slice.
func EvaluateQueue(as *[]any, is *[]int) {
	for len(*as) > 0 {
		a := Dequeue(as)
		Evaluate(a, as, is)
	}
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

///////////////////////////////////////////////////////////////////////////////////////
//                               part 4 · the operators                              //
///////////////////////////////////////////////////////////////////////////////////////

// Oper is a function that operates on a queue slice and stack slice.
type Oper func(*[]any, *[]int)

// Opers is a map of all existing Oper functions.
var Opers map[string]Oper

// part 4.1 · mathematical operators
/////////////////////////////////////

// Add2 pushes the sum of the last two integers on a stack slice.
func Add2(as *[]any, is *[]int) {
	Push(is, Pop(is)+Pop(is))
}

// Div2 pushes the quotient of the last two integers on a stack slice.
func Div2(as *[]any, is *[]int) {
	Push(is, Pop(is)/Pop(is))
}

// Mod2 pushes the modulo remainder of the last two integers on a stack slice.
func Mod2(as *[]any, is *[]int) {
	Push(is, Pop(is)%Pop(is))
}

// Mul2 pushes the product of the last two integers on a stack slice.
func Mul2(as *[]any, is *[]int) {
	Push(is, Pop(is)*Pop(is))
}

// Sub2 pushes the difference of the last two integers on a stack slice.
func Sub2(as *[]any, is *[]int) {
	Push(is, Pop(is)-Pop(is))
}

// part 4.2: stack operators
/////////////////////////////

// Dupe1 duplicates the last integer on a stack slice.
func Dupe1(as *[]any, is *[]int) {
	Push(is, Peek(is))
}

// Drop1 deletes the last integer on a stack slice.
func Drop1(as *[]any, is *[]int) {
	Pop(is)
}

// Len0 pushes the number of integers on a stack slice.
func Len0(as *[]any, is *[]int) {
	Push(is, len(*is))
}

// Swap2 swaps the last two integers on a stack slice.
func Swap2(as *[]any, is *[]int) {
	a, b := Pop(is), Pop(is)
	Push(is, a, b)
}

// part 4.3: stdin/stdout operators
////////////////////////////////////

// Input0 pushes a line from Stdin.
func Input0(as *[]any, is *[]int) {
	r := bufio.NewReader(os.Stdin)
	bs, _ := r.ReadBytes('\n')
	slices.Reverse(bs)
	for _, b := range bs {
		Push(is, int(b))
	}
}

// Print1 prints the last integer on a stack slice.
func Print1(as *[]any, is *[]int) {
	fmt.Printf("%c", Pop(is))
}

// part 4.4: logical operators
///////////////////////////////

// Def0 sets a block to a stored Oper.
func Def0(as *[]any, is *[]int) {
	xs := DequeueTo(as, "end")

	if len(xs) < 2 {
		panic(`"def" block missing name/body`)
	}

	if _, ok := xs[0].(string); !ok {
		panic(`"def" block name wrong type`)
	}

	Opers[xs[0].(string)] = func(as *[]any, is *[]int) {
		ns := make([]any, len(xs)-1)
		copy(ns, xs[1:])
		EvaluateQueue(&ns, is)
	}
}

// EvalN evaluates an entire stack slice.
func EvalN(as *[]any, is *[]int) {
	var rs []rune
	for len(*is) > 0 {
		rs = append(rs, rune(Pop(is)))
	}

	xs := Parse(string(rs))
	EvaluateQueue(&xs, is)
}

// Loop1 evaluates a block for the value of the top integer in a stack slice.
func Loop1(as *[]any, is *[]int) {
	xs := DequeueTo(as, "done")
	i := Pop(is)

	for n := 0; n < i; n++ {
		ns := make([]any, len(xs))
		copy(ns, xs)
		EvaluateQueue(&ns, is)
	}
}

// If1 evaluates a block if the last integer in a stack slice is true.
func If1(as *[]any, is *[]int) {
	xs := DequeueTo(as, "then")

	if Pop(is) != 0 {
		EvaluateQueue(&xs, is)
	}
}

// part 4.5: miscellaneous operators
/////////////////////////////////////

// Dump0 prints a stack slice to Stdout.
func Dump0(as *[]any, is *[]int) {
	var ss []string
	for _, i := range *is {
		ss = append(ss, strconv.Itoa(i))
	}

	fmt.Printf("[ %s ]\n", strings.Join(ss, " "))
}

// Nop0 does nothing.
func Nop0(as *[]any, is *[]int) {}

///////////////////////////////////////////////////////////////////////////////////////
//                             part x · the main runtime                             //
///////////////////////////////////////////////////////////////////////////////////////

func init() {
	Opers = map[string]Oper{
		// 4.1: mathematical operators
		"+": Add2,
		"/": Div2,
		"%": Mod2,
		"*": Mul2,
		"-": Sub2,

		// 4.2: stack operators
		"dupe": Dupe1,
		"drop": Drop1,
		"len":  Len0,
		"swap": Swap2,

		// 4.3: stdin/stdout operators
		"input": Input0,
		"print": Print1,

		// part 4.4: logical operators
		"def":  Def0,
		"eval": EvalN,
		"if":   If1,
		"loop": Loop1,

		// part 4.5: miscellaneous operators
		"dump": Dump0,
		"·":    Nop0,
	}
}

func main() {
	var is []int
	var as = Parse(`
		def _print  · loop print done      · end
		def _dump   · len if dump then     · end
		def _prompt · 32 62 2 _print input · end

		9 loop · _prompt eval _dump · done
		33 101 121 66 4 _print
	`)

	EvaluateQueue(&as, &is)

	// for {
	// 	fmt.Print(" ")
	// 	r := bufio.NewReader(os.Stdin)
	// 	s, _ := r.ReadString('\n')
	// 	as := Parse(s)
	// 	EvaluateQueue(&as, &is)
	// 	fmt.Printf("[ %v ]\n", is)
	// }
}
