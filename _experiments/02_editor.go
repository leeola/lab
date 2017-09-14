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
							Draw: &Draw{Chars: []rune("package")},
						},
					},
				},
			},
		},
	}

	e, _ := New(n)
	e.Start()
}
