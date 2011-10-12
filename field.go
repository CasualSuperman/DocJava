package main

import (
	//	"fmt"
	"regexp"
	"strings"
)

var fm_final, fm_private, fm_protected, fm_public, fm_static, fm_transient,
	fm_volatile Mask

var fm_masks []Mask

func init() {
	var i uint = 0
	fm_final = Mask{1 << i, "final"}
	i++
	fm_private = Mask{1 << i, "private"}
	i++
	fm_protected = Mask{1 << i, "protected"}
	i++
	fm_public = Mask{1 << i, "public"}
	i++
	fm_static = Mask{1 << i, "static"}
	i++
	fm_transient = Mask{1 << i, "transient"}
	i++
	fm_volatile = Mask{1 << i, "volatile"}
	i++
	fm_masks = append(fm_masks, fm_final, fm_private, fm_protected, fm_public,
		fm_static, fm_transient, fm_volatile)
}

// Page 196 of the Java Specification 3
// Section 8.3
type Field struct {
	fieldModifiers Maskable // Optional
	fieldType      string
	// Name in this case, can include a declaration but won't in the JavaDoc
	// context
	variableDeclarator string
	// string ";"
	javaDoc string
}

func (f Field) String() (s string) {
	s += "/**\n * "
	s += javaDoc(f.javaDoc)
	s += "\n */\n"
	s += f.fieldModifiers.String()
	s += " "
	s += f.fieldType
	s += " "
	s += f.variableDeclarator
	s += ";"
	return
}

func javaDoc(s string) string {
	s = strings.Replace(s, ">", "&gt;", -1)
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, "\n", "\n * ", -1)
	return s
}

func NewField(text string) Field {
    // Pull out name and doc, leave mods and type together
	regString := "<A NAME=\"([^\"]+)\"><!-- --></A><H3>\\n[^<]+</H3>\\n<PRE>\\n(.+)<B>[^<]+</B></PRE>\\n<DL>\\n<DD>(.+)\\n<P>"
	reg := regexp.MustCompile(regString)
	results := reg.FindStringSubmatch(text)

    // For scoping
	field_perms := ""
	field_type := ""
	line := results[2]

    // Two options
	if strings.Contains(line, "<") {
		// URL with type enclosed
		temp := strings.SplitN(line, "<", 2)
		field_perms = temp[0]
		temp_type := "<" + temp[1]
		replace := "</?A[^>]*>"
		remove := regexp.MustCompile(replace)
		result := remove.ReplaceAllString(temp_type, "")
		result = strings.Replace(result, "&gt;", ">", -1)
		field_type = strings.Replace(result, "&lt;", "<", -1)
        field_type = regexp.MustCompile("^|[^<]+\\.").ReplaceAllString(field_type, "")
	} else {
		// Builtin type
		temp_type := strings.Trim(line, " ")
		index := strings.LastIndex(temp_type, " ")
		field_perms = temp_type[:index]
		field_type = temp_type[index+1:]
	}
	field_type = strings.Trim(field_type, " ")
	mod := NewFMod(field_perms)
	return Field{mod, field_type, results[1], results[3]}
}

/* Implementing Mask Interface */

type fMod int

func NewFMod(list string) (f *fMod) {
	f = new(fMod)
	for i := 0; i < len(fm_masks); i++ {
		if strings.Contains(list, fm_masks[i].String()) {
			f.Set(fm_masks[i], true)
		}
	}
	return
}

func (f *fMod) String() (s string) {
	if f.Has(fm_public) {
		s += "public"
	} else if f.Has(fm_private) {
		s += "private"
	} else if f.Has(fm_protected) {
		s += "protected"
	}

	if f.Has(fm_final) {
		s = "final " + s
	}

	if f.Has(fm_static) {
		s += " static"
	}
	if f.Has(fm_transient) {
		s += " transient"
	}
	if f.Has(fm_volatile) {
		s += " volatile"
	}
	return
}

func (f *fMod) Has(mask Mask) bool {
	return (mask.mask & int(*f)) != 0
}

func (f *fMod) Set(mask Mask, on bool) {
	if on && !f.Has(mask) {
		*f = fMod(int(*f) ^ mask.mask)
	} else if !on && f.Has(mask) {
		*f = fMod(int(*f) | mask.mask)
	}
}
