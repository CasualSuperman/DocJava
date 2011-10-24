package main

import (
	"regexp"
	"strings"
)

var (
	classMasker FullMask = BitMasker([]Mask{
		Mask{1 << 3, "abstract"},
		Mask{1 << 4, "final"},
		Mask{1 << 5, "static"},
		Mask{1 << 6, "strictfp"}})
)

// Page 175 of the Java Specification 3
// Section 8.1
type Class struct {
	// Declaration fields
	classModifiers Maskable //Optional
	// class
	identifier string
	// <
	typeParameters []Type // Optional
	// >
	// extends
	super string // Optional
	// implements
	interfaces []string // Optional
	// {
	fieldDeclarations       []Field
	constructorDeclarations []Constructor
	methodDeclarations      []Method
	classDeclarations       []Class
	interfaceDeclarations   []Interface
	// }
	doc JavaDoc
}

func NewClass(data []string) (c Class) {
	/* 1) Modifiers
	 * 2) Class Name
	 * 3) Extends
	 * 4) Ignored
	 * 5) Implements
	 */
	preamble, nested_class, nested_interface, field, constructor, method := data[0], data[1], data[2], data[3], data[4], data[5]
	preamble_reg := regexp.MustCompile("<pre>(.*) class <span[^>]+>([^<]+)</span>\nextends ([^\\n]+)\n?(implements (.+))?</pre>")
	tempDoc := strings.Split(strings.Split(preamble, "</pre>\n")[1], "</div>\n<div class=\"summary\">")[0]
	c.doc = NewDoc(tempDoc)
	info := preamble_reg.FindStringSubmatch(preamble)
	c.classModifiers = classMasker.Apply(info[1])
	c.identifier = strings.Replace(strings.Replace(info[2], "&gt;", ">", -1), "&lt;", "<", -1)
	c.super = NewType(info[3]).String()
	if len(info[4]) > 0 {
		c.interfaces = strings.Split(RemoveUrl(info[5]), ",")
	} else {
		c.interfaces = []string{}
	}
	split := "</a>\n<ul"
	classes := strings.Split(nested_class, split)[1:]
	interfaces := strings.Split(nested_interface, split)[1:]
	fields := strings.Split(field, split)[1:]
	constructors := strings.Split(constructor, split)[1:]
	methods := strings.Split(method, split)[1:]
	for _, value := range classes {
		for _, value := range interfaces {
			split = value
		}
		split = value
	}
	for _, value := range fields {
		c.fieldDeclarations = append(c.fieldDeclarations, NewField("<a name="+value))
	}
	for _, value := range constructors {
		c.constructorDeclarations = append(c.constructorDeclarations, NewConstructor("<a name="+value))
	}
	for _, value := range methods {
		c.methodDeclarations = append(c.methodDeclarations, NewMethod("<a name="+value))
	}
	return
}

func (c Class) String() (s string) {
	s += c.doc.String()
	s += "\n"
	s += c.classModifiers.String()
	s += " class "
	s += c.identifier
	if len(c.typeParameters) > 0 {
		s += "<"
		for i, val := range c.typeParameters {
			if i > 0 {
				s += ", "
			}
			s += val.String()
		}
	}
	if c.super != "Object" {
		s += " extends "
		s += c.super
	}
	if len(c.interfaces) > 0 {
		s += " implements "
		for i, val := range c.interfaces {
			if i > 0 {
				s += ", "
			}
			s += val
		}
	}
	s += " {"
	s += "\n"
	for _, field := range c.fieldDeclarations {
		s += tab(field.String(), 1)
		s += "\n"
	}
	for _, constructor := range c.constructorDeclarations {
		s += tab(constructor.String(), 1)
		s += "\n"
	}
	for _, method := range c.methodDeclarations {
		s += tab(method.String(), 1)
		s += "\n"
	}
	s += "}"
	return
}
