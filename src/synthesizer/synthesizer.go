package synthesizer

import (
	"github.com/j-c-w/SemanticHLLVectors/src/io"
)

const MAXDEPTH = 5

// This is a simple program synthesizer that takes 
// a set of functions and tries to compose them.

func Synthesize(inputs []io.Values, outputs []io.Values) *IntFunction {
	return IntSynthesize(inputs, outputs)
}

func IntSynthesize(inputs []io.Values, outputs []io.Values) *IntFunction {
	// Try combinations of the functions
	depth := 1

	for depth < MAXDEPTH {
		// Start with just the ID function.
		var depthFunctions = IdentityFunction[:];

		for current_depth := 0; current_depth < depth; current_depth ++ {
			depthFunctions = crossProduct(depthFunctions, IntFunctions);
		}

		// Check if any of these functions satisfy the IO requirements
		for _, function := range depthFunctions {
			if IOCheck(function, inputs, outputs) {
				return &function
			}
		}

		depth ++
	}

	// Failed to find a function.
	return nil
}

func IOCheck(f IntFunction, inp []io.Values, out []io.Values) bool {
	for i := 0; i < len(inp); i ++ {
		if f.IntSim(inp[i].Value) != out[i].Value {
			return false
		}
	}

	return true
}

func crossProduct(a []IntFunction, b []Function) []IntFunction {
	result := make([]IntFunction, len(a) * len(b))

	for i := 0; i < len(a); i ++ {
		for j := 0; j < len(b); j ++ {
			result[i * j + j] = a[i].Compose(IntFunction{[]Function{b[j]}})
		}
	}

	return result
}
