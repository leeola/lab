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
	Nodes     []Node
	EndOfLine bool
}

type Keyword struct {
	*Draw
}

type Draw struct {
	Chars []rune
	Width uint8
}

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
