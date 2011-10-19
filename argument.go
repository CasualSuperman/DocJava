package main

import (
	"strings"
)

type Argument struct {
	argName string
	argType Type
}

func NewArgList(s string) []Argument {
	s = RemoveUrl(s)
	args := strings.Split(s, " ")
	result := []Argument{}
	for _, arg := range args {
		result = append(result, parse(arg))
	}
	return result
}

func parse(s string) Argument {
	halves := strings.Split(s, "&nbsp;")
	return Argument{halves[0], NewType(halves[1])}
}

func (a Argument) String() (s string) {
	s += a.argName
	s += " "
	s += a.argType.String()
	return
}
