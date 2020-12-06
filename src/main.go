package main

import (
	"fmt"
	"github.com/j-c-w/SemanticHLLVectors/src/io"
	"github.com/j-c-w/SemanticHLLVectors/src/synthesizer"
)

func main() {
	fmt.Println("Running Simple Test")
	simpleTest()
}

func simpleTest() {
	// Load all the int->int functions.  Try and convert them
	// all to each other.
	inputs := io.GenerateInts(10)

	for _, f := range Base64IntNativeFunctions {
		for _, g := range BasisExamples {
			// Let g be the basis function here.
			// Try to compute a h such that h(g(x)) = f(x)
			synthOutputs := f.ExecuteAll(inputs)
			synthInputs := g.ExecuteAll(inputs)

			// Now try to generate the synth function:
			res := synthesizer.Synthesize(synthInputs, synthOutputs)

			if res != nil {
				fmt.Println("Successful synthesizing!")
				fmt.Println("From ", f.Name)
				fmt.Println("To ", g.Name)
				fmt.Println("Using ", res.ToString())
			}
		}
	}
}
