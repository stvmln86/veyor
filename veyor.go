package main

import "github.com/stvmln86/veyor/veyor"

func main() {
	veyor.EvaluateString(veyor.Stlib+"repl", veyor.NewStack())
}
