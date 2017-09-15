package editor

type NodeType int

const (
	NodeTypeUnknown = iota
	NodeTypeClosure
	NodeTypeGroup
	NodeTypeKeyword
	NodeTypeText
)

type Node struct {
	Type    NodeType
	Closure *Closure
	Keyword *Keyword
	Text    *Text
	Group   *Group
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
	Draw
}

type Draw struct {
	Chars []rune
	Width uint8
}

// any input-able name. variable names, etc.
type Text struct {
	Draw
}

// // :=, =, { }, []
// // too many possible syntaxes, need to refine them somehow.
// type Syntax struct {
// }
//
// type File struct{}

// type Comment struct {
// }

func GroupNode(nodes ...Node) Node {
	return Node{
		Type: NodeTypeGroup,
		Group: &Group{
			Nodes: nodes,
		},
	}
}

func KeywordNode(kw string) Node {
	r := []rune(kw)
	return Node{
		Type: NodeTypeKeyword,
		Keyword: &Keyword{
			Draw: Draw{
				Chars: r,
				Width: uint8(len(r)),
			},
		},
	}
}

func TextNode() Node {
	return Node{
		Type: NodeTypeText,
		Text: &Text{},
	}
}
