package main

import (
	"regexp"
	"strings"
)

var (
	constructorMasker FullMask = BitMasker([]Mask{})
)

// Page 240 of the Java Specification 3
// Section 8.8
type Constructor struct {
	constructorModifiers Maskable // Optional
	typeParameters       string   // Optional
	// (
	formalParameterList []Argument // Optional
	// )
	// throws
	throws string // Optional
	// {
	// Body
	// }
	doc JavaDoc
}

func (c Constructor) String() (s string) {
	s += c.doc.String()
	s += "\n"
	s += c.constructorModifiers.String()
	s += " "
	s += c.typeParameters
	s += "("
	for i, param := range c.formalParameterList {
		if i > 0 {
			s += ", "
		}
		s += param.String()
	}
	s += ") {"
	s += "\n\n}"
	return
}

func NewConstructor(input string) Constructor {
	//fmt.Println(input)
	str := "<h4>([^<]+)</h4>\n<pre>([^<]+) ?(throws <a[^>]*>[^<]+</a>)?</pre>"
	//fmt.Println(str)
	data := regexp.MustCompile(str).FindStringSubmatch(input)
	//	debugPrint(data...)
	name := data[1]
	def := data[2]
	throws := ""
	mask := constructorMasker.Apply(strings.SplitN(def, "&nbsp;", 2)[0])
	docs := regexp.MustCompile("(<div[^>]+>.*</div>.*)?\n</li>\n</ul>").FindStringSubmatch(input)
	if len(data) > 3 {
		// throws clause
		throws = RemoveUrl(data[3])
	}
	//	debugPrint(doc...)
	types := NewArgList(regexp.MustCompile("\\((.*)\\)").FindStringSubmatch(def)[1])
	doc := NewDoc(docs[1])
	return Constructor{mask, name, types, throws, doc}
}
