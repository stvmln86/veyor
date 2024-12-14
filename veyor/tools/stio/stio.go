// Package stio implements standard input/output stream functions.
package stio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// Stdin is the default standard input stream.
var Stdin io.Reader = os.Stdin

// Stdout is the default standard output stream.
var Stdout io.Writer = os.Stdout

// Buffer replaces Stdin with a string buffer and Stdout with an empty buffer,
// returning Stdout.
func Buffer(s string) *bytes.Buffer {
	Stdin = bytes.NewBufferString(s)
	Stdout = bytes.NewBuffer(nil)
	return Stdout.(*bytes.Buffer)
}

// Input returns a newline-ended string from Stdin.
func Input() string {
	r := bufio.NewReader(Stdin)
	s, _ := r.ReadString('\n')
	return s
}

// Output writes a formatted string to Stdout.
func Output(s string, as ...any) {
	fmt.Fprintf(Stdout, s, as...)
}
