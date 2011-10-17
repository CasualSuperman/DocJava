package main

type Argument struct {
	argName string
	argType Type
}

func NewArgList(s string) []Argument {
	return []Argument{}
}
