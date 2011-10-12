package main

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
	constructorModifiers  Maskable // Optional
	constructorDeclarator conDeclarator
	// throws
	throws string // Optional
	// {
	// Body
	// }
}

type conDeclarator struct {
	typeParameters string // Optional
	typeTypes      Type
	// (
	formalParameterList []Argument
	// )
}

