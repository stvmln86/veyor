package veyor

// Stlib is Veyor's standard library.
const Stlib = `
	( ** Boolean Functions ** )

	def not
		( a -- a · Negate the top integer. )
		dup 2 * · swap -
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

	def drop
		( a -- · Drop the top integer. )
		if then
	end

	( ** Standard I/O Functions ** )

	def print0
		( a... -- · Print until an EOF zero. )
		loop · dup 0 eq? · if drop break else print then ·	done
	end

	def prompt
		( -- · Print a "> " prompt and push input. )
		0 32 62 print0 · input
	end
`
