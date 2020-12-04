package io

import (
	"math/rand"
)

const MAX_NUMBER=1000

func GenerateInts(number int) []Values {
	var values = make([]Values, number)

	for i := 1; i < number; i ++ {
		values[i] = Values{generateInt(), true}
	}

	return values
}

func generateInt() int {
	return rand.Intn(MAX_NUMBER)
}
