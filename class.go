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
	interfaces string // Optional
	// {
	fieldDeclarations       []Field
	constructorDeclarations []Constructor
	methodDeclarations      []Method
	classDeclarations       []Class
	interfaceDeclarations   []Interface
	// }
}

func NewClass(preamble, nested_class, nested_interface, field, constructor, method string) (c Class) {
	/* 1) Modifiers
	 * 2) Class Name
	 * 3) Extends
	 * 4) Ignored
	 * 5) Implements
	 */
	preamble_reg := regexp.MustCompile("<pre>(.*) class <span[^>]+>([^<]+)</span>\nextends ([^\\n]+)\n?(implements (.+))?</pre>")
	info := preamble_reg.FindStringSubmatch(preamble)
	c.classModifiers = NewClMod(info[1])
	c.identifier = info[2]
	c.super = NewType(info[3]).String()
	c.interfaces = RemoveUrl(info[5])
	return
}

type clMod int

func NewClMod(list string) (c *clMod) {
	c = new(clMod)
	for i := 0; i < len(cm_mods); i++ {
		if strings.Contains(list, cm_mods[i].String()) {
			c.Set(cm_mods[i], true)
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
