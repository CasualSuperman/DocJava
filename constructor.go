package main

import (
	"regexp"
	"strings"
)

var cn_public, cn_private, cn_protected Mask
var cn_masks []Mask

func init() {
	var i uint = 0
	cn_public = Mask{1 << i, "public"}
	i++
	cn_private = Mask{1 << i, "private"}
	i++
	cn_protected = Mask{1 << i, "protected"}
	i++
	cn_masks = append(cn_masks, cn_public, cn_private, cn_protected)
}

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

func NewConstructor(input string) Constructor {
	// PLACEHOLDER
	data := regexp.MustCompile("<h4>([^<]+)</h4>\\n<pre>([^<]+)</pre>(<div.+)\n</li>\n</ul>").FindStringSubmatch(input)
	name := data[1]
	def := data[2]
	mod := NewCMod(strings.SplitN(def, "&nbsp;", 2)[1])
	doc := NewDoc(data[3])
	types := NewArgList(regexp.MustCompile("\\((.*)\\)").FindStringSubmatch(def)[1])
	return Constructor{mod, name, types, "", doc}
}

type cMod int

func NewCMod(list string) (c *cMod) {
	c = new(cMod)
	for i := 0; i < len(cn_masks); i++ {
		if strings.Contains(list, cn_masks[i].String()) {
			c.Set(cn_masks[i], true)
		}
	}
	return
}

func (c *cMod) String() (s string) {
	if c.Has(cn_public) {
		s += "public"
	} else if c.Has(cn_protected) {
		s += "protected"
	} else if c.Has(cn_private) {
		s += "private"
	}
	return
}

func (c *cMod) Has(mask Mask) bool {
	return (mask.mask & int(*c)) != 0
}

func (c *cMod) Set(mask Mask, on bool) {
	if on && !c.Has(mask) {
		*c = cMod(int(*c) ^ mask.mask)
	} else if !on && c.Has(mask) {
		*c = cMod(int(*c) | mask.mask)
	}
}
