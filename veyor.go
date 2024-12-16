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

// BoolToInt returns an integer from a boolean.
func BoolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}

// Try panics on a true boolean with a formatted error message.
func Try(b bool, s string, as ...any) {
	if b {
		panic(fmt.Sprintf(s, as...))
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                         part three · collection functions                         //
///////////////////////////////////////////////////////////////////////////////////////

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
		a := (*as)[0]
		(*as) = (*as)[1:]
		switch a := a.(type) {
		case int:
			Push(is, a)
		case string:
			f, ok := Opers[a]
			Try(!ok, "invalid operator %q", a)
			f(as, is)
		}
	}
}

// EvaluateCopy evaluates a copy of a queue against a stack.
func EvaluateCopy(as []any, is *[]int) {
	xs := make([]any, len(as))
	copy(xs, as)
	Evaluate(&xs, is)
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

// init initialises the Opers map.
func init() {
	Opers = map[string]Oper{
		// math operators
		"+": func(as *[]any, is *[]int) { Push(is, Pop(is)+Pop(is)) },
		"/": func(as *[]any, is *[]int) { Push(is, Pop(is)/Pop(is)) },
		"%": func(as *[]any, is *[]int) { Push(is, Pop(is)%Pop(is)) },
		"*": func(as *[]any, is *[]int) { Push(is, Pop(is)*Pop(is)) },
		"-": func(as *[]any, is *[]int) { Push(is, Pop(is)-Pop(is)) },
		">": func(as *[]any, is *[]int) { Push(is, BoolToInt(Pop(is) > Pop(is))) },

		// stack operators
		"dup":  func(as *[]any, is *[]int) { Push(is, Peek(is)) },
		"len":  func(as *[]any, is *[]int) { Push(is, len(*is)) },
		"swap": func(as *[]any, is *[]int) { Push(is, Pop(is), Pop(is)) },
		"roll": func(as *[]any, is *[]int) {
			a, b, c := Pop(is), Pop(is), Pop(is)
			Push(is, b, a, c)
		},

		// block operators
		"break": func(as *[]any, is *[]int) { Break = true },
		"(":     func(as *[]any, is *[]int) { DequeueTo(as, ")") },

		"assert": func(as *[]any, _ *[]int) {
			as1 := DequeueTo(as, "=>")
			is1, is2 := new([]int), new([]int)
			for _, a := range DequeueTo(as, "end") {
				*is2 = append(*is2, a.(int))
			}

			EvaluateCopy(as1, is1)
			Try(!slices.Equal(*is1, *is2),
				"assert: %v should equal %v, not %v\n", as1, *is2, *is1,
			)
		},

		"def": func(as *[]any, is *[]int) {
			xs := DequeueTo(as, "end")
			_, ok := xs[0].(string)
			Try(len(xs) < 2 || !ok, `invalid def block`)

			Opers[xs[0].(string)] = func(as *[]any, is *[]int) {
				EvaluateCopy(xs[1:], is)
			}
		},

		"if": func(as *[]any, is *[]int) {
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
		},

		"loop": func(as *[]any, is *[]int) {
			xs := DequeueTo(as, "done")

			for {
				EvaluateCopy(xs, is)
				if Break {
					Break = false
					break
				}
			}
		},

		// i/o & eval operators
		"dump":  func(as *[]any, is *[]int) { fmt.Fprintf(Stdout, ": %v\n", *is) },
		"print": func(as *[]any, is *[]int) { fmt.Fprintf(Stdout, "%c", Pop(is)) },

		"eval": func(as *[]any, is *[]int) {
			var rs []rune
			for len(*is) > 0 {
				if i := Pop(is); i == 0 {
					break
				} else {
					rs = append(rs, rune(i))
				}
			}

			EvaluateString(string(rs), is)
		},

		"input": func(as *[]any, is *[]int) {
			r := bufio.NewReader(Stdin)
			bs, _ := r.ReadBytes('\n')
			slices.Reverse(bs)

			Push(is, 0)
			for _, b := range bs {
				Push(is, int(b))
			}
		},
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                          part six · the standard library                          //
///////////////////////////////////////////////////////////////////////////////////////

// Stlib is Veyor's standard library.
const Stlib = `
	def ·      ( --       )                    end
	def drop   ( a --     ) if then            end
	def eq?    ( a b -- b ) - if 0 else 1 then end
	def even?  ( a -- b   ) 2 swap % false?    end
	def false? ( a -- b   ) if 0 else 1 then   end
	def neq?   ( a b -- c ) eq? false?         end
	def odd?   ( a -- b   ) 2 swap % true?     end
	def true?  ( a -- b   ) if 1 else 0 then   end

	def clear ( ... --   )
		loop · len false? if break else drop then · done
	end

	def repl ( -- )
		loop · 0 32 62 print0 input eval len if dump then · done
	end

	def print0 ( a... -- )
		loop · dup false? if drop break else print then · done
	end
`

///////////////////////////////////////////////////////////////////////////////////////
//                           part seven · the main runtime                           //
///////////////////////////////////////////////////////////////////////////////////////

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
