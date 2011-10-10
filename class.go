package main

type ClassMask Mask

const (
    abstract  ClassMask = 1 << (iota + 3)
    final
    static
    strictfp
)

// Page 175 of the Java Specification 3
// Section 8.1
type Class struct {
    // Declaration fields
    classModifiers byte //Optional
    // string "Class"
    identifier string
    typeParameters Type
    super *Class // Optional
    interfaces []*Interface // Optional
    // Body
    //
}

