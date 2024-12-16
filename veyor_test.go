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

func TestOperators(t *testing.T) {
	// setup
	b := mockStreams("t\n")

	// success
	evalCode(`
		( ** Math Operators ** )
		assert 1 2 + => 3 end
		assert 3 6 / => 2 end
		assert 3 5 % => 2 end
		assert 2 3 * => 6 end
		assert 3 5 - => 2 end

		( ** Stack Operators ** )
		assert 1 dup     => 1 1   end
		assert 1 len     => 1 1   end
		assert 1 2 swap  => 2 1   end
		assert 1 2 3 rot => 2 3 1 end

		( ** Block Operators ** )
		assert ( )                =>   end
		assert def x 1 end x      => 1 end
		assert 1 if 1 then        => 1 end
		assert 0 if 1 else 2 then => 2 end
		assert loop 1 break done  => 1 end

		( ** I/O & Eval Operators ** )
		assert 1 dump    => 1 end
		assert 0 49 eval => 1 end
		assert input     => 0 10 116 end
		116 print
	`)

	assert.Equal(t, ": [1]\nt", b.String())
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
		assert 2 not => -2 end

		( ** Conditional Functions ** )
		assert 1 1 eq?   · 1 2 eq?   => 1 0 end
		assert 1   even? · 2   even? => 0 1 end
		assert 1 1 neq?  · 1 2 neq?  => 0 1 end
		assert 1   odd?  · 2   odd?  => 1 0 end
		assert 0   zero? · 1 zero?   => 1 0 end

		( ** Miscellaneous Functions ** )
		assert · => end

		( ** Stack Functions ** )
		assert 1 2 3 clear => end
		assert 1 2 3 drop  => 1 2 end

		( ** Standard I/O Functions ** )
		assert 0 10 116 115 101 116 print0 => end
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
