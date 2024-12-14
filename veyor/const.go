package veyor

import (
	"io"
	"os"
)

// Break indicates that the current loop should exit.
var Break bool = true

// Opers is a map of all defined operator functions.
// var Opers map[string]func(Queue, Stack)

// Stdin is the default input Reader.
var Stdin io.Reader = os.Stdin

// Stdout is the default output Writer.
var Stdout io.Writer = os.Stdout
