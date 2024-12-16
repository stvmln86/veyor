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
	EvaluateString(Stlib+s, new([]int))
}

func mockStreams(s string) *bytes.Buffer {
	Stdin = bytes.NewBufferString(s)
	Stdout = bytes.NewBuffer(nil)
	return Stdout.(*bytes.Buffer)
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part two · helper functions                            //
///////////////////////////////////////////////////////////////////////////////////////

func TestTry(t *testing.T) {
	// success
	Try(false, "%s", "panic")
}

///////////////////////////////////////////////////////////////////////////////////////
//                         part three · collection functions                         //
///////////////////////////////////////////////////////////////////////////////////////

func TestDequeue(t *testing.T) {
	// setup
	as := &[]any{1}

	// success
	a := Dequeue(as)
	assert.Equal(t, 1, a)
	assert.Empty(t, *as)

	// error - queue is empty
	defer func() { assert.Equal(t, "queue is empty", recover()) }()
	Dequeue(as)
}

func TestDequeueTo(t *testing.T) {
	// setup
	as := &[]any{1, "end"}

	// success
	xs := DequeueTo(as, "end")
	assert.Equal(t, []any{1}, xs)
	assert.Empty(t, *as)

	// error - queue is missing end
	defer func() { assert.Equal(t, "queue is missing end", recover()) }()
	DequeueTo(as, "end")
}

func TestPeek(t *testing.T) {
	// success
	i := Peek(&[]int{1})
	assert.Equal(t, 1, i)
}

func TestPop(t *testing.T) {
	// setup
	is := &[]int{1}

	// success
	i := Pop(is)
	assert.Equal(t, 1, i)
	assert.Empty(t, *is)

	// error - stack is empty
	defer func() { assert.Equal(t, "stack is empty", recover()) }()
	Pop(is)
}

func TestPush(t *testing.T) {
	// setup
	is := new([]int)

	// success
	Push(is, 1, 2)
	assert.Equal(t, []int{1, 2}, *is)
}

///////////////////////////////////////////////////////////////////////////////////////
//                    part four · parsing and evaluating functions                   //
///////////////////////////////////////////////////////////////////////////////////////

func TestEvaluate(t *testing.T) {
	// setup
	as := &[]any{1, 2, "+"}
	is := new([]int)

	// success
	Evaluate(as, is)
	assert.Equal(t, []int{3}, *is)
}

func TestEvaluateString(t *testing.T) {
	// setup
	is := new([]int)

	// success
	EvaluateString("1 2 +", is)
	assert.Equal(t, []int{3}, *is)
}

func TestParse(t *testing.T) {
	// success
	as := Parse("1 2 3 a b c\n")
	assert.Equal(t, []any{1, 2, 3, "a", "b", "c"}, as)
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
		1 2 3 rot    · assert 2 3 1 end
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

	assert.Equal(t, ": [123]\nt", b.String())
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
		1 1 eq? 1 2 eq?          · assert 1 0   end
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
