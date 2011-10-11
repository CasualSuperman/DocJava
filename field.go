package main

import (
    "fmt"
    "regexp"
)

var fm_final, fm_private, fm_protected, fm_public, fm_static, fm_transient,
	fm_volatile Mask

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
}

// Page 196 of the Java Specification 3
// Section 8.3
type Field struct {
	fieldModifiers Maskable // Optional
	fieldType      Type
	// Name in this case, can include a declaration but won't in the JavaDoc
	// context
	variableDeclarator string
	// string ";"
}

func (f Field) String() (s string) {
	s += f.fieldModifiers.String()
	s += f.fieldType.String()
	s += ";"
	return
}

func NewField(text string) Field {
	/*
	<A NAME="head"><!-- --></A><H3>
	head</H3>
	<PRE>
	private <A HREF="../../../../../edu/wcu/cs/cs363/project02/BoundedLinkedSet.ListNode.html" title="class in edu.wcu.cs.cs363.project02">BoundedLinkedSet.ListNode</A> <B>head</B></PRE>
	<DL>
	<DD>The start of the linked list.
	<P>
	<DL>
	</DL>
	</DL>
	*/
	regString := "<A NAME=\"([^\"]+)\"><!-- --></A><H3>\\n[^<]+</H3>\\n<PRE>\\n(([^<]*)<A HREF=\"[^\"]+\" title=\"[^\"]+\">([^<]+)</A>|([^ ]+)([^ ]+)) <B>[^<]+</B></PRE>\\n<DL>\\n<DD>([^\\n]+)"
	fmt.Println(regString)
	reg := regexp.MustCompile(regString)
	results := reg.FindStringSubmatch(text)
    for i := 0; i < len(results); i++ {
        fmt.Println("{")
        fmt.Println("\t\"",results[i], "\"")
        fmt.Println("}")
    }
	/*
	 * 1) Field name
	 * 2) Modifiers
	 * 3) Type
	 * 4) Description
	 */
	visibilityString := "public|private|protected"
	vis := regexp.MustCompile(visibilityString)
	visibility := vis.FindString(results[2])
	fmt.Println(visibility)
    return Field{}
}

/* Implementing Mask Interface */

type fMod int

func (f fMod) String() (s string) {
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

func (f fMod) Has(mask Mask) bool {
	return (mask.mask & int(f)) != 0
}
