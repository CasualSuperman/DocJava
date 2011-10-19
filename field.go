package main

import (
	"regexp"
	"strings"
)

var (
	fm_masks []Mask = []Mask{
		Mask{1 << 3, "static"},
		Mask{1 << 4, "final"},
		Mask{1 << 5, "transient"},
		Mask{1 << 6, "volatile"}}
)

// Page 196 of the Java Specification 3
// Section 8.3
type Field struct {
	fieldModifiers Maskable // Optional
	fieldType      Type
	// Name in this case, can include a declaration but won't in the JavaDoc
	// context
	variableDeclarator string
	// string ";"
	javaDoc JavaDoc
}

func (f Field) String() (s string) {
	s += f.javaDoc.String()
	s += "\n"
	s += f.fieldModifiers.String()
	s += " "
	s += f.fieldType.String()
	s += " "
	s += f.variableDeclarator
	s += ";"
	return
}

func NewField(text string) Field {
	// Pull out name and doc, leave mods and type together
	regString := "<pre>([^&]+)&nbsp;(.+) ([^ ]+)</pre>\n(<div[^>]*>.*</div>.*)\n</li>\n</ul>"
	reg := regexp.MustCompile(regString)
	results := reg.FindStringSubmatch(text)

	mods := results[1]
	uType := results[2]
	sType := NewType(uType)
	name := strings.Replace(results[3], " ", "", -1)
	docs := NewDoc(results[4])
	mod := NewFMod(mods)
	return Field{mod, sType, name, docs}
}

/* Implementing Mask Interface */

type fMod struct {
	base_mask
}

func NewFMod(list string) (f fMod) {
	f.base_mask = NewBaseMask(list)
	for _, mask := range fm_masks {
		if strings.Contains(list, mask.String()) {
			f.Set(mask, true)
		}
	}
	return
}

func (f fMod) String() (s string) {
	s = f.base_mask.String()
	for _, mask := range fm_masks {
		if f.Has(mask) {
			if s != "" {
				s += " "
			}
			s += mask.String()
		}
	}
	return
}
