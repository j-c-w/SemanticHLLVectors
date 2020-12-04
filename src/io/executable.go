package io

import (
	"os/exec"
)

type Executable interface {
	Execute(Values) Values
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

// Implementation of a native (i.e. go function) program.
type NativeExecutable struct {
	f func(int) int
}


func (n NativeExecutable) Execute(inputs Values) Values {
	in := inputs.Value

	return Values{n.f(in), true}
}
