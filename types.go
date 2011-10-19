package main

import (
	"strings"
	"regexp"
)

// Page 184 of the Java Specification 3
// Section 8.1.4
type Type struct {
	typeDeclSpecifier string
	typeArguments     []Type // Optional
}

func (t Type) String() (s string) {
	s += t.typeDeclSpecifier
	if len(t.typeArguments) > 0 {
		s += "<"
		s += t.typeArguments[0].String()
		for i := 1; i < len(t.typeArguments); i++ {
			s += ", "
			s += t.typeArguments[i].String()
		}
		s += ">"
	}
	return
}

func NewType(usType string) (t Type) {
	sType := ""
	// If not a basic type
	if strings.Contains(usType, "<") {
		// URL with type enclosed
		usType = RemoveUrl(usType)
		usType = strings.Replace(usType, "&gt;", ">", -1)
		usType = strings.Replace(usType, "&lt;", "<", -1)
		sType = regexp.MustCompile(", |,|^|[^<,]+ ?\\.").ReplaceAllString(usType, "")
	} else {
		sType = strings.Trim(usType, " ")
	}
	t.typeDeclSpecifier = sType
	return
}
