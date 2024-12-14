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

	"github.com/stvmln86/veyor/veyor"
)

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

	Push(is, 0)
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

// Eq2 pushes 1 or 0 if the top two integers are equal.
func Eq2(as *[]any, is *[]int) {
	if Pop(is) == Pop(is) {
		Push(is, 1)
	} else {
		Push(is, 0)
	}
}

// EvalN evaluates an entire stack slice.
func EvalN(as *[]any, is *[]int) {
	var rs []rune
	for len(*is) > 0 {
		i := Pop(is)
		if i == 0 {
			break
		} else {
			rs = append(rs, rune(i))
		}
	}

	if s := strings.TrimSpace(string(rs)); s != "" {
		xs := Parse(string(rs))
		EvaluateQueue(&xs, is)
	}
}

// Loop1 evaluates a block for the value of the top integer in a stack slice.
func Loop1(as *[]any, is *[]int) {
	xs := DequeueTo(as, "done")

	for range Pop(is) {
		ns := make([]any, len(xs))
		copy(ns, xs)
		EvaluateQueue(&ns, is)

		if Break {
			Break = false
			break
		}
	}
}

// If1 evaluates a block if the last integer in a stack slice is true.
func If1(as *[]any, is *[]int) {
	xs := DequeueTo(as, "then")

	if Pop(is) != 0 {
		EvaluateQueue(&xs, is)
	}
}

// Not1 inverts an integer.
func Not1(as *[]any, is *[]int) {
	if Pop(is) == 0 {
		Push(is, 1)
	} else {
		Push(is, 0)
	}
}

// part 4.5: miscellaneous operators
/////////////////////////////////////

// Break0 sets a special flag to break a loop.
func Break0(as *[]any, is *[]int) {
	Break = true
}

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

func main() {
	veyor.EvaluateString(`
		def print0 · len loop · print · dupe 0 eq? if drop break then · done · end
		def prompt · 0 32 62 print0 input · end

		9 loop · prompt eval · len if dump then · done
		0 33 101 121 66 print0
	`, veyor.NewStack())
}
