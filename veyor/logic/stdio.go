package logic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/stvmln86/veyor/veyor/atoms/cell"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

// Stdin is the default standard input stream.
var Stdin io.Reader = os.Stdin

// Stdout is the default standard output stream.
var Stdout io.Writer = os.Stdout

// Input0 pushes a newline-ended string from Stdin as Cells.
func Input0(_ *queue.Queue, st *stack.Stack) error {
	var cs []cell.Cell
	s, _ := bufio.NewReader(Stdin).ReadString('\n')
	for _, r := range s {
		cs = append(cs, cell.Cell(r))
	}

	slices.Reverse(cs)
	st.PushAll(cs)
	return nil
}

// PrintN write all Cells from a Stack until a newline to Stdout.
func PrintN(_q *queue.Queue, s *stack.Stack) error {
	for !s.Empty() {
		c, err := s.Pop()
		if err != nil {
			return err
		}

		fmt.Fprintf(Stdout, "%c", c)
		if c == cell.Cell(10) {
			break
		}
	}

	return nil
}
