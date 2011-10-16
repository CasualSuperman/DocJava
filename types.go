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

func NewType(uType string) (t Type) {
	sType := ""
	// If not a basic type
	if strings.Contains(uType, "<") {
		// URL with type enclosed
		replace := "</?a[^>]*>"
		remove := regexp.MustCompile(replace)
		uType = remove.ReplaceAllString(uType, "")
		uType = strings.Replace(uType, "&gt;", ">", -1)
		uType = strings.Replace(uType, "&lt;", "<", -1)
		sType = regexp.MustCompile("^|[^<]+\\.").ReplaceAllString(uType, "")
		/*	} else {
			// Builtin type
			sType := strings.Trim(uType, " ")
		*/
	} else {
		sType = strings.Replace(uType, " ", "", -1)
	}
	t.typeDeclSpecifier = sType
	return
}
