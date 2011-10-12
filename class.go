package main

var cm_abstract, cm_final, cm_private, cm_protected, cm_public, cm_static,
	cm_strictfp Mask

var cm_mods []Mask

func init() {
	var i uint = 0
	cm_abstract = Mask{1 << i, "abstract"}
	i++
	cm_final = Mask{1 << i, "final"}
	i++
	cm_private = Mask{1 << i, "private"}
	i++
	cm_protected = Mask{1 << i, "protected"}
	i++
	cm_public = Mask{1 << i, "public"}
	i++
	cm_static = Mask{1 << i, "static"}
	i++
	cm_strictfp = Mask{1 << i, "strictfp"}
	i++
	cm_mods = append(cm_mods, cm_abstract, cm_final, cm_private, cm_protected, cm_public, cm_static, cm_strictfp)
}

// Page 175 of the Java Specification 3
// Section 8.1
type Class struct {
	// Declaration fields
	classModifiers byte //Optional
	// string "Class"
	identifier     string
	typeParameters Type
	super          string   // Optional
	interfaces     []string // Optional
	// Body
	// []*ClassBodyDeclaration
	fieldDeclarations       []Field
	constructorDeclarations []Constructor
	methodDeclarations      []Method
	classDeclarations       []Class
	interfaceDeclarations   []Interface
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
