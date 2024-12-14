package main

import "github.com/stvmln86/veyor/veyor"

func main() {
	veyor.EvaluateString(`
		def print0
			( Print all integers up to an EOF zero. )
			loop · print · dup 0 eq? · if drop break then · done
		end

		def prompt
			( Print a "> " prompt and push input. )
			0 32 62 print0 · input
		end

		def repl
			( Run the read-eval-print loop. )
			loop · prompt eval · len if dump then · done
		end

		repl · 0 33 101 121 66 print0
	`, veyor.NewStack())
}
