package main

import . "github.com/leeola/lab/editor"

func main() {
	n := Node{
		Type: NodeTypeClosure,
		Closure: &Closure{
			Options: Options{
				{
					Desc: "package",
					Node: Node{
						Type: NodeTypeKeyword,
						Keyword: &Keyword{
							String: "package",
						},
					},
				},
			},
		},
	}

	e, _ := New(n)
	e.Start()
}
