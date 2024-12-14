package stio

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	// success
	b := Buffer("Input.\n")
	s, _ := bufio.NewReader(Stdin).ReadString('\n')
	fmt.Fprint(Stdout, "Output.\n")
	assert.Equal(t, "Input.\n", s)
	assert.Equal(t, "Output.\n", b.String())
}

func TestInput(t *testing.T) {
	// setup
	Buffer("Input.\n")

	// success
	s := Input()
	assert.Equal(t, "Input.\n", s)
}

func TestOutput(t *testing.T) {
	// setup
	b := Buffer("")

	// success
	Output("%s\n", "Output.")
	assert.Equal(t, "Output.\n", b.String())
}
