package main

//Type
type Type struct {
	typeDeclSpecifier string
	typeArguments     []*string // Optional
}

func (t Type) String() (s string) {
	s += t.typeDeclSpecifier
	if len(t.typeArguments) > 0 {
		s += "<"
		s += *t.typeArguments[0]
		for i := 1; i < len(t.typeArguments); i++ {
			s += ", "
			s += *t.typeArguments[i]
		}
		s += ">"
	}
	return
}
