///////////////////////////////////////////////////////////////////////////////////////
//                veyor · a minimal scheme-ish language in Go · v0.0.0               //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func atomise(s string) any {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	return s
}

func tokenise(s string) []string {
	s = strings.Replace(s, "[", " [ ", -1)
	s = strings.Replace(s, "]", " ] ", -1)
	return strings.Fields(s)
}

func expressionise(ss *[]string) any {
	if len(*ss) == 0 {
		panic("unexpected EOF")
	}

	s := (*ss)[0]
	*ss = (*ss)[1:]
	fmt.Printf("%q %#v\n", s, ss)

	switch s {
	case "[":
		var as []any
		for (*ss)[0] != "]" {
			as = append(as, expressionise(ss))
		}

		*ss = (*ss)[1:]
		return as

	case "]":
		panic("unexpected closing paren")

	default:
		return atomise(s)
	}
}

///////////////////////////////////////////////////////////////////////////////////////

func main() {
	ss := tokenise(`[def double [x] [+ x x]]`)
	fmt.Printf("%#v\n", expressionise(&ss))
}
