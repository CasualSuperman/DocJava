package main

import (
	"strings"
)

type Argument struct {
	argName string
	argType Type
}

func NewArgList(s string) []Argument {
	if strings.Trim(s, " ") == "" {
		return []Argument{}
	}
	s = RemoveUrl(s)
	args := strings.Split(s, " ")
	result := []Argument{}
	for _, arg := range args {
		result = append(result, parseArgs(arg))
	}
	return result
}

func parseArgs(s string) Argument {
	halves := strings.Split(s, "&nbsp;")
	return Argument{halves[0], NewType(halves[1])}
}

func (a Argument) String() (s string) {
	s += a.argName
	s += " "
	s += a.argType.String()
	return
}
