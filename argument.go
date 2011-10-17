package main

type Argument struct {
	argName string
	argType Type
}

func NewArgList(s string) []Argument {
	return []Argument{}
}

func (a Argument) String() (s string) {
	s += a.argName
	s += " "
	s += a.argType.String()
	return
}
