///////////////////////////////////////////////////////////////////////////////////////
//                veyor · a minimal stack language in Go · unit tests                //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////////////
//                   part one · unit testing variables & functions                   //
///////////////////////////////////////////////////////////////////////////////////////

func evalCode(s string) {
	EvaluateString(Stlib+s, NewStack())
}

func mockStreams(s string) *bytes.Buffer {
	Stdin = bytes.NewBufferString(s)
	Stdout = bytes.NewBuffer(nil)
	return Stdout.(*bytes.Buffer)
}

///////////////////////////////////////////////////////////////////////////////////////
//                             part two · the queue type                             //
///////////////////////////////////////////////////////////////////////////////////////

func TestNewQueue(t *testing.T) {
	// success
	q := NewQueue(123)
	assert.Equal(t, []any{123}, q.Atoms)
}

func TestQueueDequeue(t *testing.T) {
	// setup
	q := NewQueue(123)

	// success
	a := q.Dequeue()
	assert.Equal(t, 123, a)
	assert.Empty(t, q.Atoms)

	// error - Queue is empty
	defer func() { assert.Equal(t, "Queue is empty", recover()) }()
	q.Dequeue()
}

func TestQueueDequeueTo(t *testing.T) {
	// setup
	q := NewQueue(123, "end")

	// success
	as := q.DequeueTo("end")
	assert.Equal(t, []any{123}, as)
	assert.Empty(t, q.Atoms)

	// error - Queue is missing end
	defer func() { assert.Equal(t, "Queue is missing end", recover()) }()
	q.DequeueTo("end")
}

func TestQueueEmpty(t *testing.T) {
	// success - true
	ok := NewQueue().Empty()
	assert.True(t, ok)

	// success - false
	ok = NewQueue(123).Empty()
	assert.False(t, ok)
}

func TestQueueEnqueue(t *testing.T) {
	// setup
	q := NewQueue()

	// success
	q.Enqueue(123, 456)
	assert.Equal(t, []any{123, 456}, q.Atoms)
}

func TestQueueIndex(t *testing.T) {
	// success
	i := NewQueue(123, 456).Index(456)
	assert.Equal(t, 1, i)
}

func TestQueueLen(t *testing.T) {
	// success
	i := NewQueue(123).Len()
	assert.Equal(t, 1, i)
}

func TestQueueString(t *testing.T) {
	// success
	s := NewQueue(123, "abc").String()
	assert.Equal(t, "123 abc", s)
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part three · the stack type                            //
///////////////////////////////////////////////////////////////////////////////////////

func TestNewStack(t *testing.T) {
	// success
	s := NewStack(123)
	assert.Equal(t, []int{123}, s.Items)
}

func TestStackEmpty(t *testing.T) {
	// success - true
	ok := NewStack().Empty()
	assert.True(t, ok)

	// success - false
	ok = NewStack(123).Empty()
	assert.False(t, ok)
}

func TestStackLen(t *testing.T) {
	// success
	i := NewStack(123).Len()
	assert.Equal(t, 1, i)
}

func TestStackPeek(t *testing.T) {
	// success
	i := NewStack(123).Peek()
	assert.Equal(t, 123, i)
}

func TestStackPop(t *testing.T) {
	// setup
	s := NewStack(123)

	// success
	i := s.Pop()
	assert.Equal(t, 123, i)
	assert.Empty(t, s.Items)

	// error - Stack is empty
	defer func() { assert.Equal(t, "Stack is empty", recover()) }()
	s.Pop()
}

func TestStackPush(t *testing.T) {
	// setup
	s := NewStack()

	// success
	s.Push(123, 456)
	assert.Equal(t, []int{123, 456}, s.Items)
}

func TestStackString(t *testing.T) {
	// success
	s := NewStack(123, 456).String()
	assert.Equal(t, "123 456", s)
}

///////////////////////////////////////////////////////////////////////////////////////
//                    part four · parsing and evaluating functions                   //
///////////////////////////////////////////////////////////////////////////////////////

func TestEvaluate(t *testing.T) {
	// setup
	q := NewQueue(1, 2, "+")
	s := NewStack()

	// success
	Evaluate(q, s)
	assert.Equal(t, []int{3}, s.Items)
}

func TestEvaluateSlice(t *testing.T) {
	// setup
	s := NewStack()

	// success
	EvaluateSlice([]any{1, 2, "+"}, s)
	assert.Equal(t, []int{3}, s.Items)
}

func TestEvaluateString(t *testing.T) {
	// setup
	s := NewStack()

	// success
	EvaluateString("1 2 +", s)
	assert.Equal(t, []int{3}, s.Items)
}

func TestParse(t *testing.T) {
	// success
	as := Parse("123 abc\n")
	assert.Equal(t, []any{123, "abc"}, as)
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part five · operator functions                          //
///////////////////////////////////////////////////////////////////////////////////////

func TestMathOperators(t *testing.T) {
	// success
	evalCode(`
		1 2 + · assert 3 end
		1 2 + · assert 3 end
		3 6 / · assert 2 end
		3 5 % · assert 2 end
		2 3 * · assert 6 end
		3 5 - · assert 2 end
	`)
}

func TestStackOperators(t *testing.T) {
	// success
	evalCode(`
		123 dup      · assert 123 123 end
		123 len      · assert 123 1   end
		123 456 swap · assert 456 123 end
	`)
}

func TestLogicalOperators(t *testing.T) {
	// success
	evalCode(`
		1 2 eq? · assert 0 end
		2 2 eq? · assert 1 end
	`)
}

func TestBlockOperators(t *testing.T) {
	// success
	evalCode(`
		123 ( comment )        · assert 123 end
		def x 123 end · x      · assert 123 end
		1 if 123 then          · assert 123 end
		0 if 123 else 456 then · assert 456 end
		loop 123 break done    · assert 123 end
	`)
}

func TestIOEvalOperators(t *testing.T) {
	// setup
	b := mockStreams("test\n")

	// success
	evalCode(`
		123 dump        · assert 123 end
		0 51 50 49 eval · assert 123 end
		input · assert 0 10 116 115 101 116 end
		116 print
	`)

	assert.Equal(t, "[ 123 ]\nt", b.String())
}

///////////////////////////////////////////////////////////////////////////////////////
//                          part six · the standard library                          //
///////////////////////////////////////////////////////////////////////////////////////

func TestStlib(t *testing.T) {
	// setup
	b := mockStreams("test\n")

	// success
	evalCode(`
		( ** Boolean Functions ** )
		2 not · assert -2 end

		( ** Conditional Functions ** )
		1 even? 2 even?          · assert 0 1   end
		1 1 neq? 1 2 neq?        · assert 0 1   end
		1 odd?  2 odd?           · assert 1 0   end
		-1 zero? 0 zero? 1 zero? · assert 0 1 0 end

		( ** Miscellaneous Functions ** )
		· assert end

		( ** Stack Functions ** )
		1 2 3 clear · assert end
		1 2 3 drop  · assert 1 2 end

		( ** Standard I/O Functions ** )
		0 10 116 115 101 116 print0 · assert end
	`)

	assert.Equal(t, "test\n", b.String())
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part seven · the main runtime                           //
///////////////////////////////////////////////////////////////////////////////////////

func TestInit(t *testing.T) {
	// success
	assert.NotEmpty(t, Opers)
}
