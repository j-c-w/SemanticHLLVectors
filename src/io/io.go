package io

import (
	"os/exec"
)

func get_outputs(inputs Values, executable Executable) Values {
	// We expect the executable to have a self-timeout function.
	out,err := exec.Command(executable.name, inputs.toString()).Output()

	if err != nil {
		// Invalid values
		return Values {0, false};
	}

	return valuesFromString(string(out))
}
