package synthesizer

import (
	"strings"
)

type IntFunction struct {
	functions []Function
}

type Function struct {
	name string
	id int
}

// This script contains a number of functoins used
// as a basis for synthesis.
var IntFunctions = []Function {
	Function {"zero", 0},
	Function {"timestwo", 1},
	Function {"half", 2},
	Function {"increment", 3},
	Function {"decrement", 4},
	Function {"id", 5},
}

var IdentityFunction = []IntFunction {
	IntFunction{ []Function{ Function {"id", 5 }} },
}

func id(x int) int {
	return x
}

func zero(x int) int {
	return 0
}

func timestwo(x int) int {
	return x * 2
}

func half(x int) int {
	return x / 2
}

func increment(x int) int {
	return x + 1
}

func decrement(x int) int {
	return x - 1
}

func runIntFunction(f Function, x int) int {
	switch f.id {
	case 0:
		return zero(x)
	case 1:
		return timestwo(x)
	case 2:
		return half(x)
	case 3:
		return increment(x)
	case 4:
		return decrement(x)
	case 5:
		return id(x)
	}

	panic("Unknown function ")
}

func (f IntFunction) IntSim(x int) int {
	// Innermost to outermost order.
	for i := len(f.functions) - 1; i >= 0; i -- {
		x = runIntFunction(f.functions[i], x)
	}

	return x
}

func (f IntFunction) Compose (g IntFunction) IntFunction {
	result := IntFunction { append(f.functions, g.functions...) }
	return result
}

func (f IntFunction) toString() string {
	var builder strings.Builder

	for i, fun := range f.functions {
		builder.WriteString(fun.name)
		if i != len(f.functions) - 1 {
			builder.WriteString(" o ")
		}
	}

	return builder.String()
}
