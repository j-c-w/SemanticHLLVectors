package io

import (
	"strconv"
)

type Values struct {
	Value int
	Valid bool
}

func (value Values) toString() string {
	return strconv.Itoa(value.Value)
}

func valuesFromString(str string) Values {
	res, err := strconv.Atoi(str)
	if err != nil {
		return Values {res, true}
	} else {
		return Values {res, false}
	}
}
