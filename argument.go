package main

import (
	"regexp"
	"strings"
)

type Argument struct {
	argName string
	argType Type
}

func NewArgList(s string) []Argument {
	s = regexp.MustCompile("</?a[^>]*>").ReplaceAllString(s, "")
	args := strings.Split(s, " ")
	result := []Argument{}
	for i := 0; i < len(args); i++ {
		result = append(result, parse(args[i]))
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
