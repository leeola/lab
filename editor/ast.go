package editor

type NodeType int

const (
	NodeTypeUnknown = iota
	NodeTypeClosure
	NodeTypeGroup
	NodeTypeKeyword
)

type Node struct {
	Type    NodeType
	Closure *Closure
	Keyword *Keyword
	// Text    *Text
	// Group   *Group
}

type Options []*Option

type Closure struct {
	Nodes   []Node
	Options []*Option
}

type Option struct {
	Desc string
	Node Node
}

type Group struct {
	String string
}

type Keyword struct {
	String string
}

// type Option struct {
// }
//
// // any input-able name. variable names, etc.
// type Text struct {
// 	Pos
// }
//
// // :=, =, { }, []
// // too many possible syntaxes, need to refine them somehow.
// type Syntax struct {
// }
//
// type File struct{}
//
// type Pos struct {
// 	Line uint
// 	X    uint
// }
//
// type Comment struct {
// }
//
// func GroupNode(v Group) Node {
// 	return Node{
// 		Type:  NodeTypeGroup,
// 		Group: &v,
// 	}
// }
//
// func KeywordNode(v Keyword) Node {
// 	return Node{
// 		Type:    NodeTypeKeyword,
// 		Keyword: &v,
// 	}
// }
//
// func TextNode(v Text) Node {
// 	return Node{
// 		Type:    NodeTypeText,
// 		Keyword: &v,
// 	}
// }
