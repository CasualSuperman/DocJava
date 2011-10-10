package main

type Visibility string

const (
    private Visibility = iota
    protected
    public
)

type Doc string
type Type string
type Variable struct{
    Name string
    Type Type
}
type Field struct{
    Variable
    Visibility
    Doc
}
type Argument struct{
    Variable
    Doc
}
type Exception struct{
    Name, Reason string
}
type Function struct{
    Name string
    Arguments []Argument
    Return Type
    Description, Precondition, Postcondition, Returns string
    Exceptions []Exception
}
type Constructor Function
type Class struct{
    Name, Description, Note, Version string
    Authors []string
    Methods []Function
    Constructors []Constructor
    Fields []Variable
    SubClasses []Class
}


func (t *Type) Init(name string) *Type{
    *t = Type(name)
    return t
}

func (v *Variable) Init(name, typeName string) *Variable{
    v.Type = *new(Type).Init(typeName)
    v.Name = name
    return v
}

func (f *Field) Init(Var Variable, mods, doc string) *Field{
    f.Variable = Var
    f.Visibility = mods
    f.Doc = doc
    return f
}

func (a *Argument) Init(Var Variable, doc string) *Argument{
    a.Variable = Var
    a.Doc = doc
    return a
}

func (e *Exception) Init(name, reason string) *Exception{
    e.Name = name
    e.Reason = reason
    return e
}

func (f *Function) Init(name string, returns Type) *Function{
    f.Name = name
    f.Return = returns
    return f
}

func (f *Function) AddArgument(Arg Argument){
    f.Arguments = append(f.Arguments, Arg)
}

func (f *Function) AddException(Ex Exception){
    f.Exceptions = append(f.Exceptions, Ex)
}

func (f *Function) AddPrecondition(pre string){
    f.Precondition = pre
}

func (f *Function) AddPostcondition(post string){
    f.Postcondition = post
}

func (f *Function) AddReturnDoc(doc string){
    f.Returns = doc
}

func (f *Function) AddDescription(doc string){
    f.Description = doc
}

func (c *Constructor) Init() *Constructor{
    return c
}
/*
func (c *Class) Init(name string) *Class{
    c.Name = name
    return c
}*/
func (c *Class) AddMethod(Func Function){
    c.Methods = append(c.Methods, Func)
}
func (c *Class) AddField(Var Variable){
    c.Fields = append(c.Fields, Var)
}
func (c *Class) AddConstructor(Func Constructor){
    c.Constructors = append(c.Constructors, Func)
}
func (c *Class) AddAuthor(author string){
    c.Authors = append(c.Authors, author)
}
func (c *Class) AddNote(note string){
    c.Note = note
}
func (c *Class) AddVersion(ver string){
    c.Version = ver
}
func (c *Class) AddDescription(desc string){
    c.Description = desc
}
func (c *Class) AddNode(note string){
    c.Note = note
}
