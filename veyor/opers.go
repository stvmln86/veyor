package veyor

// Oper is a function that operates on a Queue and Stack.
type Oper func(*Queue, *Stack)

// Opers is a map of all defined operator functions.
var Opers map[string]Oper

// init initialises the Opers map.
func init() {
	Opers = map[string]Oper{
		// opmaths.go
		"+": OpAdd,
		"/": OpDivide,
		"%": OpModulo,
		"*": OpMultiply,
		"-": OpSubtract,
	}
}
