package main

type FieldMask Mask

const (
    fm_final FieldMask = 1 << (iota + 3)
    fm_private
    fm_protected
    fm_public
    fm_static
    fm_transient
    fm_volatile
)

// Page 196 of the Java Specification 3
// Section 8.3
type Field struct {
    fieldModifiers fMod // Optional
    fieldType Type
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

type fMod byte

func (f fMod) String() (s string) {
    mask := f
    if (mask.has(fm_public)) {
        s += "public"
    } else if (mask.has(fm_private)) {
        s += "private"
    } else if (mask.has(fm_protected)) {
        s += "protected"
    }

    if (mask.has(fm_final)) {
        s = "final " + s
    }

    if (mask.has(fm_static)) {
        s += " static"
    }
    if (mask.has(fm_transient)) {
        s += " transient"
    }
    if (mask.has(fm_volatile)) {
        s += " volatile"
    }
    return
}

func (f fMod) has(mask FieldMask) bool {
    return (int(mask) & int(f)) == 0
}
