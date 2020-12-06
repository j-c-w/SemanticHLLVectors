package main

// Contains a number of native tests

import (
	"github.com/j-c-w/SemanticHLLVectors/src/io"
)

// Int-> Int snippets taken from https://github.com/golang/go/blob/master/src/encoding/base64/base64.go
// Encode
func f1(x int) int {
	return x / 3 * 3
}

func f2(x int) int {
	return x - x % 3
}

func f3(x int) int {
	return x << 16 | x << 8 | x
}

func f4(x int) int {
	x = x / 3 * 4
	if x < 4 {
		x = 4
	}

	return x
}

func f5(x int) int {
	return x / 4 * 3
}

func f6(x int) int {
	return x << 26 | x << 20 | x << 14 | x << 8
}

func f7(x int) int {
	return x * 6 / 8
}

type IntFunc func(int) int

var Base64IntFunctions = []IntFunc { f1, f2, f3, f4, f5, f6, f7 };
var Base64IntNativeFunctions = []io.NativeExecutable {
	io.NativeExecutable{f1, "f1"},
	io.NativeExecutable{f2, "f2"},
	io.NativeExecutable{f3, "f3"},
	io.NativeExecutable{f4, "f4"},
	io.NativeExecutable{f5, "f5"},
	io.NativeExecutable{f6, "f6"},
	io.NativeExecutable{f7, "f7"},
}

// Hand-crafted examples
func h1(x int) int {
	return x * 4
}

func h2(x int) int {
	return x
}

func h3(x int) int {
	return 2 * x + 2
}

func h4(x int) int {
	return x / 2 * 2
}

var HandcraftedNativeFunctions = []io.NativeExecutable {
	io.NativeExecutable{h1, "h1"},
	io.NativeExecutable{h2, "h2"},
	io.NativeExecutable{h3, "h3"},
	io.NativeExecutable{h4, "h4"},
}

// Basis function examples
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

// Int basis examples
var BasisExamples = []io.NativeExecutable {
	io.NativeExecutable{id, "id"},
	io.NativeExecutable{zero, "zero"},
	io.NativeExecutable{timestwo, "timestwo"},
	io.NativeExecutable{half, "half"},
	io.NativeExecutable{increment, "increment"},
	io.NativeExecutable{decrement, "decrement"},
}
