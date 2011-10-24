package main

import (
	"strings"
)

type Argument struct {
	argType Type
	argName string
}

func NewArgList(s string) []Argument {
	if strings.Trim(s, " ") == "" {
		return []Argument{}
	}
	s = RemoveUrl(s)
	args := strings.Split(s, ",")
	result := []Argument{}
	for _, arg := range args {
		result = append(result, parseArgs(arg))
	}
	return result
}

func parseArgs(s string) Argument {
	halves := strings.Split(s, "&nbsp;")
	return Argument{NewType(halves[0]), halves[1]}
}

func (a Argument) String() (s string) {
	s += a.argType.String()
	s += " "
	s += a.argName
	return
}
