package editor

type Node struct {
	Type    int
	Keyword *Keyword
	Name    *Name
}

type Keyword struct {
	Pos
	String string
}

type Type struct {
	Pos
	String string
	Name   *Name
}

// any input-able name. variable names, etc.
type Name struct {
	Pos
}

// :=, =, { }, []
// too many possible syntaxes, need to refine them somehow.
type Syntax struct {
}

type Group struct{}
type Closure struct{}
type File struct{}

type Pos struct {
	Line uint
	X    uint
}
