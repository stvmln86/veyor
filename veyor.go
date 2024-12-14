package main

import (
	"flag"

	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
	"github.com/stvmln86/veyor/veyor/logic"
	"github.com/stvmln86/veyor/veyor/tools/stio"
)

var (
	mStack = stack.New()
	mQueue = queue.New()
)

func eval(s string) error {
	as, err := logic.Parse(s)
	if err != nil {
		stio.Output("Error: %s.\n", err.Error())
		return err
	}

	mQueue.EnqueueAll(as)
	if err := logic.EvaluateQueue(mQueue, mStack); err != nil {
		stio.Output("Error: %s.\n", err.Error())
		return err
	}

	if !mStack.Empty() {
		stio.Output("[ %s ]\n", mStack.String())
	}

	return nil
}

func main() {
	fComm := flag.String("c", "", "execute string and exit")
	flag.Parse()

	switch {
	case *fComm != "":
		eval(*fComm)

	default:
		for {
			stio.Output("> ")
			if err := eval(stio.Input()); err != nil {
				continue
			}
		}
	}
}
