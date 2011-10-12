package main


type Constructor struct {
    constructorModifiers Maskable // Optional
    constructorDeclarator conDeclarator
    // throws
    throws string // Optional
    // {
    // Body
    // }
}

type conDeclarator struct {
    typeParameters string // Optional
    typeTypes Type
    // (
    formalParameterList []Argument
    // )
}

var cn_public, cn_private, cn_protected Mask
