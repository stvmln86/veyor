package logic

import (
	"fmt"

	"github.com/stvmln86/veyor/veyor/atoms/word"
	"github.com/stvmln86/veyor/veyor/langs/queue"
	"github.com/stvmln86/veyor/veyor/langs/stack"
)

// Def0 sets an enclosed block to a stored Oper.
func Def0(q *queue.Queue, _ *stack.Stack) error {
	as, err := q.DequeueTo(word.Word("end"))
	if err != nil {
		return fmt.Errorf(`"def" block missing "end"`)
	}

	if len(as) < 2 {
		return fmt.Errorf(`"def" block missing name/body`)
	}

	if _, ok := as[0].(word.Word); !ok {
		return fmt.Errorf(`"def" block name wrong type`)
	}

	Opers[word.Word(as[0].String())] = func(_ *queue.Queue, s *stack.Stack) error {
		return EvaluateQueue(queue.New(as[1:]...), s)
	}

	return nil
}

// If1 evaluates an enclosed block if the top Cell is true.
func If1(q *queue.Queue, s *stack.Stack) error {
	as, err := q.DequeueTo(word.Word("then"))
	if err != nil {
		return fmt.Errorf(`"if" block missing "then"`)
	}

	c, err := s.Pop()
	if err != nil {
		return err
	}

	if c.Bool() {
		return EvaluateQueue(queue.New(as...), s)
	}

	return nil
}
