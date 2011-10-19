package main

import (
	"regexp"
	"strings"
)

var cm_abstract, cm_final, cm_private, cm_protected, cm_public, cm_static,
	cm_strictfp Mask

var cm_mods []Mask

func init() {
	var i uint = 0
	cm_abstract = Mask{1 << i, "abstract"}
	i++
	cm_final = Mask{1 << i, "final"}
	i++
	cm_private = Mask{1 << i, "private"}
	i++
	cm_protected = Mask{1 << i, "protected"}
	i++
	cm_public = Mask{1 << i, "public"}
	i++
	cm_static = Mask{1 << i, "static"}
	i++
	cm_strictfp = Mask{1 << i, "strictfp"}
	i++
	cm_mods = append(cm_mods, cm_abstract, cm_final, cm_private, cm_protected, cm_public, cm_static, cm_strictfp)
}

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

func NewClass(preamble, nested_class, nested_interface, field, constructor, method string) (c Class) {
	/* 1) Modifiers
	 * 2) Class Name
	 * 3) Extends
	 * 4) Ignored
	 * 5) Implements
	 */
	preamble_reg := regexp.MustCompile("<pre>(.*) class <span[^>]+>([^<]+)</span>\nextends ([^\\n]+)\n?(implements (.+))?</pre>")
	tempDoc := strings.Split(strings.Split(preamble, "</pre>\n")[1], "</div>\n<div class=\"summary\">")[0]
	c.doc = NewDoc(tempDoc)
	info := preamble_reg.FindStringSubmatch(preamble)
	c.classModifiers = NewClMod(info[1])
	c.identifier = strings.Replace(strings.Replace(info[2], "&gt;", ">", -1), "&lt;", "<", -1)
	c.super = NewType(info[3]).String()
	if len(info[4]) > 0 {
		c.interfaces = strings.Split(RemoveUrl(info[5]), ",")
	} else {
		c.interfaces = []string{}
	}
	split := "<a name="
	//classes := strings.Split(nested_class, split)
	//interfaces := strings.Split(nested_interface, split)
	fields := strings.Split(field, split)[1:]
	constructors := strings.Split(constructor, split)[1:]
	methods := strings.Split(method, split)[1:]
	for _, value := range fields {
		c.fieldDeclarations = append(c.fieldDeclarations, NewField("<a name=" + value))
	}
	for _, value := range constructors {
		c.constructorDeclarations = append(c.constructorDeclarations, NewConstructor("<a name=" + value))
	}
	for _, value := range methods {
		c.methodDeclarations = append(c.methodDeclarations, NewMethod("<a name=" + value))
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

type clMod int

func NewClMod(list string) (c *clMod) {
	c = new(clMod)
	for _, mod := range cm_mods {
		if strings.Contains(list, mod.String()) {
			c.Set(mod, true)
		}
	}
	return
}

func (c *clMod) String() (s string) {
	if c.Has(cm_public) {
		s = "public"
	} else if c.Has(cm_private) {
		s = "private"
	} else if c.Has(cm_protected) {
		s = "protected"
	}
	if c.Has(cm_final) {
		s += " final"
	}
	if c.Has(cm_abstract) {
		s += " abstract"
	}
	if c.Has(cm_static) {
		s += " static"
	}
	if c.Has(cm_strictfp) {
		s += " strictfp"
	}
	return
}

func (c *clMod) Has(mask Mask) bool {
	return (mask.mask & int(*c)) != 0
}

func (c *clMod) Set(mask Mask, on bool) {
	if on && !c.Has(mask) {
		*c = clMod(int(*c) ^ mask.mask)
	} else if !on && c.Has(mask) {
		*c = clMod(int(*c) ^ mask.mask)
	}
}
