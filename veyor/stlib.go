package veyor

// Stlib is Veyor's standard library.
const Stlib = `
	( ** Boolean Functions ** )

	def not
		( a -- a · Negate the top integer. )
		dup 2 * · swap -
	end

	( ** Conditional Functions ** )

	def even?
		( a -- a · Push 1 if the top integer is even. )
		2 swap % · 0 eq? · if 1 else 0 then
	end

	def odd?
		( a -- a · Push 1 if the top integer is odd. )
		2 swap % · 0 eq? · if 0 else 1 then
	end

	def zero?
		( a -- a · Push 1 if the top integer is zero. )
		0 eq? · if 1 else 0 then
	end

	( ** Interactive Functions ** )

	def repl
		( -- · Launch a read-eval-print loop. )
		loop · prompt eval · len if dump then · done
	end

	( ** Miscellaneous Functions ** )

	def ·
		( -- · Do nothing. )
	end

	( ** Stack Functions ** )

	def clear
		( a... -- · Drop all Stack integers. )
		loop · len zero? · if break else drop · then · done
	end

	def drop
		( a -- · Drop the top integer. )
		if then
	end

	( ** Standard I/O Functions ** )

	def print0
		( a... -- · Print until an EOF zero. )
		loop · dup zero? if drop break else print then · done
	end

	def prompt
		( -- · Print a "> " prompt and push input. )
		0 32 62 print0 · input
	end
`
