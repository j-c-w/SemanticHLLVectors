package synthesizer

import (
	"github.com/j-c-w/SemanticHLLVectors/src/io"
	"fmt"
)

const MAXDEPTH = 4
const DEBUG = false
var Consts = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// This is a simple program synthesizer that takes 
// a set of functions and tries to compose them.

func Synthesize(inputs []io.Values, outputs []io.Values) Node {
	return IntSynthesize(inputs, outputs)
}

func IntSynthesize(inputs []io.Values, outputs []io.Values) Node {
	// Try combinations of the functions
	depth := MAXDEPTH

	// Generate functions with 'depth' operators.
	depthFunctions := make(chan Node)
	go GenerateFunctionsWithDepth(depth, Consts, depthFunctions)

	// Check if any of these functions satisfy the IO requirements
	for function := range depthFunctions {
		if DEBUG {
			fmt.Println("Checking")
			fmt.Println(function.ToString())
		}
		if IOCheck(function, inputs, outputs) {
			return function
		}
	}

	// Failed to find a function.
	return nil
}

func GenerateFunctionsWithDepth(depth int, consts []int, c chan Node) {
	if depth == 0 {
		for _, con := range consts {
			c <- IntChild{con}
		}
		// TODO -- more than one variable for more than
		// one variable functions :)
		c <- VariableChild{"x"}
		close(c)
		return
	}

	// For each depth, we add every operation:
	smallerSubTrees := make(chan Node)
	go GenerateFunctionsWithDepth(depth / 2, consts, smallerSubTrees)
	// Repor tthe smple trees first
	for tree := range smallerSubTrees {
		c <- tree
	}

	// Need to recreate the subtrees
	subTrees1 := make(chan Node)
	go GenerateFunctionsWithDepth(depth / 2, consts, subTrees1)

	// Now, add the parent trees with every combination
	// of subtrees.
	for subTree1 := range subTrees1 {
		subTrees2 := make(chan Node)
		go GenerateFunctionsWithDepth(depth / 2, consts, subTrees2)
		for subTree2 := range subTrees2 {
			for _, op := range binaryOperators {
				nextNode := BinaryNode{op, subTree1, subTree2}
				if !nextNode.IsSimplifiable() {
					c <- nextNode
				}
			}
		}

		// Also do the unary operators.
		for _, op := range unaryOperators {
			nextNode := UnaryNode{op, subTree1}
			if !nextNode.IsSimplifiable() {
				c <- nextNode
			}
		}
	}

	close(c)
}

func IOCheck(f Node, inp []io.Values, out []io.Values) bool {
	if len(inp) != len(out) {
		panic( "Expect input/output lengths to be equal")
	}
	for i := 0; i < len(inp); i ++ {
		if f.Evaluate(inp[i].Value) != out[i].Value {
			return false
		}
	}

	return true
}
