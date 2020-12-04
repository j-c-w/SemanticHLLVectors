package io

import (
	"strconv"
)

type Values struct {
	value int
	valid bool
}

func (value Values) toString() string {
	return strconv.Itoa(value.value)
}

func valuesFromString(str string) Values {
	res, err := strconv.Atoi(str)
	if err != nil {
		return Values {res, true}
	} else {
		return Values {res, false}
	}
}
