package main

import "github.com/stvmln86/veyor/veyor"

func main() {
	s := veyor.NewStack()
	veyor.EvaluateString(veyor.Stlib+"repl", s)
}
