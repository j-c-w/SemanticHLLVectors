package io

import (
	"os/exec"
)

type Executable interface {
	Execute(Values) Values
	ExecuteAll([]Values) []Values
}

// Implementation of a program that may correspond e.g.
// to a different language.
type Program struct {
	name string
}

func NewProgram(name string) Program {
	return Program { name }
}

func (e Program) ToString () string {
	return e.name
}

func (p Program) Execute(inputs Values) Values {
	// We expect the executable to have a self-timeout function.
	out,err := exec.Command(p.name, inputs.toString()).Output()

	if err != nil {
		// Invalid values
		return Values {0, false};
	}

	return valuesFromString(string(out))
}

func (e Program) ExecuteAll(values []Values) []Values {
	outputs := make([]Values, len(values))
	for i, value := range values {
		outputs[i] = e.Execute(value)
	}

	return outputs
}

// Implementation of a native (i.e. go function) program.
type NativeExecutable struct {
	F func(int) int
	Name string
}


func (n NativeExecutable) Execute(inputs Values) Values {
	in := inputs.Value

	return Values{n.F(in), true}
}

func (n NativeExecutable) ExecuteAll(values []Values) []Values {
	outputs := make([]Values, len(values))

	for i, value := range values {
		outputs[i] = n.Execute(value)
	}

	return outputs
}
