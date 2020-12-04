package io

import (
)

func GetOutputs(inputs []Values, executable Executable) []Values {
	var results = make([]Values, len(inputs))

	for i := 0; i < len(inputs); i ++ {
		output := executable.Execute(inputs[i])
		results[i] = output
	}

	return results
}
