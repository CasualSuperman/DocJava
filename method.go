package main

import (
	"regexp"
	"strings"
)

var me_public, me_private, me_protected, me_final, me_static, me_abstract, me_strictfp, me_synchronized, me_native Mask
var me_masks []Mask

func init() {
	var i uint = 0
	me_public = Mask{1 << i, "public"}
	i++
	me_protected = Mask{1 << i, "protected"}
	i++
	me_private = Mask{1 << i, "private"}
	i++
	me_final = Mask{1 << i, "final"}
	i++
	me_static = Mask{1 << i, "static"}
	i++
	me_abstract = Mask{1 << i, "abstract"}
	i++
	me_strictfp = Mask{1 << i, "strictfp"}
	i++
	me_synchronized = Mask{1 << i, "synchronized"}
	i++
	me_native = Mask{1 << i, "native"}
	i++
	me_masks = []Mask{me_public, me_protected, me_private, me_final, me_static, me_abstract, me_strictfp, me_synchronized, me_native}
}

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
	mMask := NewMMod(data[1])
	mType := NewType(data[2])
	mName := data[3]
	mArgs := NewArgList(data[4])
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

type mMod int

func NewMMod(list string) (m *mMod) {
	m = new(mMod)
	for _, mask := range me_masks {
		if strings.Contains(list, mask.String()) {
			m.Set(mask, true)
		}
	}
	return
}

func (m *mMod) String() (s string) {
	if m.Has(me_public) {
		s += "public"
	} else if m.Has(me_private) {
		s += "private"
	} else if m.Has(me_protected) {
		s += "protected"
	}
	return
}

func (m *mMod) Has(mask Mask) bool {
	return (mask.mask & int(*m)) != 0
}

func (m *mMod) Set(mask Mask, on bool) {
	if on && !m.Has(mask) {
		*m = mMod(int(*m) ^ mask.mask)
	} else if !on && m.Has(mask) {
		*m = mMod(int(*m) ^ mask.mask)
	}
}
