package io

type Executable struct {
	name string
}

func NewExecutable(name string) Executable {
	return Executable { name }
}

func (e Executable) ToString () string {
	return e.name
}
