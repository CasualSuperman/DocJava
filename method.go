package main

import (
	"regexp"
	"strings"
)

var (
	methodMasker FullMask = BitMasker([]Mask{
		Mask{1 << 3, "final"},
		Mask{1 << 4, "static"},
		Mask{1 << 5, "abstract"},
		Mask{1 << 6, "strictfp"},
		Mask{1 << 7, "synchronized"},
		Mask{1 << 8, "native"}})
)

type Method struct {
	methodModifiers Maskable // Optional
	methodType      Type     // Optional
	methodName      string
	// (
	formalParameterList []Argument
	// )
	// throws
	throws string
	// {
	// Body
	// }
	doc JavaDoc
}

func NewMethod(s string) (m Method) {
	/* 1) Method modifiers
	 * 2) Method type
	 * 3) Method name
	 * 4) Arguments
	 * 5) (Optional) throws
	 */
	str := "<pre>(.+)&nbsp;(.+)&nbsp;(.+)\\((.*)\\)[^<]*(throws <a[^>]*>.+</a>)?</pre>"
	data := regexp.MustCompile(str).FindStringSubmatch(s)
	mMask := methodMasker.Apply(data[1])
	mType := NewType(data[2])
	mName := data[3]
	mArgs := NewArgList(regexp.MustCompile("\n[ \t]*").ReplaceAllString(data[4], ""))
	mThrow := data[5]
	mDoc := NewDoc("<div" + strings.SplitN(s, "<div", 2)[1])
	return Method{mMask, mType, mName, mArgs, mThrow, mDoc}
}

func (m Method) String() (s string) {
	s += m.doc.String()
	s += "\n"
	s += m.methodModifiers.String()
	s += " "
	s += m.methodType.String()
	s += " "
	s += m.methodName
	s += "("
	for i, param := range m.formalParameterList {
		if i > 0 {
			s += ", "
		}
		s += param.String()
	}
	s += ")"
	if m.throws != "" {
		s += " "
	}
	s += m.throws
	s += " {"
	s += "\n\n}"
	return
}
