package main

var im_abstract, im_public, im_private, im_protected, im_static, im_strictfp Mask
var im_mods []Mask

func init() {
	var i uint = 0
	im_abstract = Mask{1 << i, "abstract"}
	i++
	im_public = Mask{1 << i, "public"}
	i++
	im_private = Mask{1 << i, "private"}
	i++
	im_protected = Mask{1 << i, "protected"}
	i++
	im_static = Mask{1 << i, "static"}
	i++
	im_strictfp = Mask{1 << i, "strictfp"}
	i++
	im_mods = append(im_mods, im_abstract, im_public, im_private, im_protected, im_static, im_strictfp)
}

type Interface struct {
	iModifiers byte //Optional
	// interface
	iIdentifier     string
	iTypeParameters Type     // Optional
	iExtends        []string // Optional
	// {
	// Body
	// }
}
