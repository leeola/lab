package main

import . "github.com/leeola/lab/editor"

func main() {
	n := Node{
		Type: NodeTypeClosure,
		Closure: &Closure{
			Options: Options{
				{
					Desc: "package",
					Node: GroupNode(
						KeywordNode("package"),
						TextNode(),
					),
				},
			},
		},
	}

	e, _ := New(n)
	e.Start()
}
