package main

var (
	im_mods []Mask = []Mask{
		Mask{1 << 3, "abstract"},
		Mask{1 << 4, "static"},
		Mask{1 << 5, "strictfp"}}
)

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
