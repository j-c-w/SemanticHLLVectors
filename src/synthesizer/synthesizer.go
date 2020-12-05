package synthesizer

import (
	"github.com/j-c-w/SemanticHLLVectors/src/io"
	"fmt"
)

const MAXDEPTH = 5
var Consts = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// This is a simple program synthesizer that takes 
// a set of functions and tries to compose them.

func Synthesize(inputs []io.Values, outputs []io.Values) Node {
	return IntSynthesize(inputs, outputs)
}

func IntSynthesize(inputs []io.Values, outputs []io.Values) Node {
	// Try combinations of the functions
	depth := 1

	for depth < MAXDEPTH {
		// Generate functions with 'depth' operators.
		depthFunctions := GenerateFunctionsWithDepth(depth, Consts)

		// Check if any of these functions satisfy the IO requirements
		for _, function := range depthFunctions {
			if IOCheck(function, inputs, outputs) {
				return function
			}
		}

		depth ++
	}

	// Failed to find a function.
	return nil
}

func GenerateFunctionsWithDepth(depth int, consts []int) []Node {
	if depth == 0 {
		var results []Node = make([]Node, len(consts) + 1)
		for i, c := range consts {
			results[i] = IntChild{c}
		}
		// TODO -- more than one variable for more than
		// one variable functions :)
		results[len(results) - 1] = VariableChild{"x"}
		return results
	}

	// For each depth, we add every operation:
	subTrees := GenerateFunctionsWithDepth(depth / 2, consts)

	// Now, add the parent trees with every combination
	// of subtrees.
	results := make([]Node, len(subTrees) * len(subTrees) * len(operators))
	for i, subTree1 := range subTrees {
		for j, subTree2 := range subTrees {
			for k, op := range operators {
				results[i + len(subTrees) * (j + len(subTrees) * k)] = BinaryNode{op, subTree1, subTree2}
			}
		}
	}

	for i, r := range results {
		if r == nil {
			fmt.Println(len(subTrees))
			fmt.Println(len(operators))
			fmt.Println(i)
			fmt.Println("FAILED")
			panic("FAILED")
		}
	}

	return results
}

func IOCheck(f Node, inp []io.Values, out []io.Values) bool {
	if len(inp) != len(out) {
		panic( "Expect input/output lengths to be equal")
	}
	for i := 0; i < len(inp); i ++ {
		fmt.Println(inp[i])
		fmt.Println(out[i])
		fmt.Println(i)
		if f.Evaluate(inp[i].Value) != out[i].Value {
			return false
		}
	}

	return true
}
