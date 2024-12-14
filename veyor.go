package main

import "github.com/stvmln86/veyor/veyor"

func main() {
	veyor.EvaluateString(`
		def print0 · len loop · print · dupe 0 eq? if drop break then · done · end
		def prompt · 0 32 62 print0 input · end

		9 loop · prompt eval · len if dump then · done
		0 33 101 121 66 print0
	`, veyor.NewStack())
}
