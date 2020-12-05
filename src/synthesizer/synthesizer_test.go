package synthesizer

import (
	"testing"
	"github.com/j-c-w/SemanticHLLVectors/src/io"
	"fmt"
)

func TestGen(t *testing.T) {
	inputs := [3]io.Values{io.Values{1, true}, io.Values{2, true}, io.Values{3, true}};
	outputs := [3] io.Values{io.Values{1, true}, io.Values{4, true}, io.Values{7, true}};

	fun := Synthesize(inputs[:], outputs[:])
	fmt.Println(fun.ToString())
}

func TestNext(t *testing.T) {
	inputs := [2]io.Values{io.Values{1, true}, io.Values{2, true}};
	outputs := [2] io.Values{io.Values{2, true}, io.Values{5, true}};

	fun := Synthesize(inputs[:], outputs[:])
	fmt.Println(fun.ToString())
}
