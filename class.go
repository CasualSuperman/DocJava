package main

type ClassMask Mask

const (
    cm_abstract  ClassMask = 1 << iota
    cm_final
    cm_private
    cm_protected
    cm_public
    cm_static
    cm_strictfp
)

// Page 175 of the Java Specification 3
// Section 8.1
type Class struct {
    // Declaration fields
    classModifiers byte //Optional
    // string "Class"
    identifier string
    typeParameters Type
    super string // Optional
    interfaces []*Interface // Optional
    // Body
    // []*ClassBodyDeclaration
    fieldDeclarations []*Field
    constructorDeclarations []*Constructor
    methodDeclarations []*Method
    classDeclarations []*Class
    interfaceDeclarations []*Interface
}

/* ClassBodyDeclaration
 * can be one of:
 *   ClassMemberDeclaration
 *   InstanceInitializer
 *   StaticInitializer
 *   ConstructorDeclaration
 */

/* ClassMemberDeclaration
 * can be one of:
 *   FieldDeclaration
 *   MethodDeclaration
 *   ClassDeclaration
 *   InterfaceDeclaration
 *   ;
 */
