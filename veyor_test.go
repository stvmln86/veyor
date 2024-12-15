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

func assertCode(t *testing.T, s string, is ...int) *Stack {
	st := NewStack()
	EvaluateString(Stlib+s, st)

	if len(is) == 0 {
		assert.Empty(t, st.Items)
	} else {
		assert.Equal(t, is, st.Items)
	}

	return st
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

func TestQueueContains(t *testing.T) {
	// success - true
	ok := NewQueue(123).Contains(123)
	assert.True(t, ok)

	// success - false
	ok = NewQueue().Contains(123)
	assert.False(t, ok)
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

	// error - Queue is insufficient
	defer func() { assert.Equal(t, "Queue is insufficient", recover()) }()
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

// part five-one · math operators
//////////////////////////////////

func TestOpAdd(t *testing.T) {
	// success
	assertCode(t, "1 2 +", 3)
}

func TestOpDivide(t *testing.T) {
	// success
	assertCode(t, "3 6 /", 2)
}

func TestOpModulo(t *testing.T) {
	// success
	assertCode(t, "3 5 %", 2)
}

func TestOpMultiply(t *testing.T) {
	// success
	assertCode(t, "2 3 *", 6)
}

func TestOpSubtract(t *testing.T) {
	// success
	assertCode(t, "3 5 -", 2)
}

// part five-two · stack operators
///////////////////////////////////

func TestOpDup(t *testing.T) {
	// success
	assertCode(t, "123 dup", 123, 123)
}

func TestOpLen(t *testing.T) {
	// success
	assertCode(t, "123 len", 123, 1)
}

func TestOpSwap(t *testing.T) {
	// success
	assertCode(t, "123 456 swap", 456, 123)
}

// part five-three · logical operators
///////////////////////////////////////

func TestOpEq(t *testing.T) {
	// success
	assertCode(t, "1 2 eq?", 0)
	assertCode(t, "1 1 eq?", 1)
}

// part five-four · block operators
/////////////////////////////////////

func TestOpBreak(t *testing.T) {
	// success
	assertCode(t, "break")
	assert.True(t, Break)
}

func TestOpComment(t *testing.T) {
	// success
	assertCode(t, "( comment )")
}

func TestOpDef(t *testing.T) {
	// success
	assertCode(t, "def foo 123 end foo", 123)

	// error - "def" missing name/body
	defer func() { assert.Equal(t, `"def" missing name/body`, recover()) }()
	assertCode(t, "def end")

	// error - "def" name is wrong type
	defer func() { assert.Equal(t, `"def" name is wrong type`, recover()) }()
	assertCode(t, "def 123 456 end")
}

func TestOpIf(t *testing.T) {
	// success - single true
	assertCode(t, "1 if 123 then", 123)

	// success - single false
	assertCode(t, "0 if 123 then")

	// success - else true
	assertCode(t, "1 if 123 else 456 then", 123)

	// success - else false
	assertCode(t, "0 if 123 else 456 then", 456)
}

func TestOpLoop(t *testing.T) {
	// success
	assertCode(t, "loop 123 break done", 123)
}

// part five-five · i/o & eval operators
/////////////////////////////////////////

func TestOpDump(t *testing.T) {
	// setup
	b := mockStreams("")

	// success
	assertCode(t, "1 2 3 dump", 1, 2, 3)
	assert.Equal(t, "[ 1 2 3 ]\n", b.String())
}

func TestOpEval(t *testing.T) {
	// success
	assertCode(t, "0 43 32 50 32 49 eval", 3)
}

func TestOpInput(t *testing.T) {
	// setup
	mockStreams("test\n")

	// success
	assertCode(t, "input", 0, 10, 116, 115, 101, 116)
}

func TestOpPrint(t *testing.T) {
	// setup
	b := mockStreams("")

	// success
	assertCode(t, "116 print")
	assert.Equal(t, "t", b.String())
}

///////////////////////////////////////////////////////////////////////////////////////
//                          part six · the standard library                          //
///////////////////////////////////////////////////////////////////////////////////////

func TestStlib(t *testing.T) {
	// setup
	b := mockStreams("test\n")

	// success - boolean functions
	assertCode(t, "2 not", -2)

	// success - conditional functions
	assertCode(t, "1 even? 2 even?", 0, 1)
	assertCode(t, "1 odd? 2 odd?", 1, 0)
	assertCode(t, "0 zero? 1 zero?", 1, 0)

	// success - miscellaneous functions
	assertCode(t, "·")

	// success - stack functions
	assertCode(t, "1 2 3 clear")
	assertCode(t, "1 2 drop", 1)

	// success - standard i/o functions
	assertCode(t, "0 10 116 115 101 116 print0")
	assert.Equal(t, "test\n", b.String())
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part seven · the main runtime                           //
///////////////////////////////////////////////////////////////////////////////////////

func TestInit(t *testing.T) {
	// success
	assert.NotEmpty(t, Opers)
}
