package main

import . "github.com/leeola/lab/editor"

func main() {
  c := Closure{
	Nodes: []Node{
  	Group{
		Nodes: []Node{
			KeywordNode({
				String: "package",
			}),
			NameNode({
			}),
			NameNode({
  			       Start: `"`,
			}),
		},
	},
	},
  }

fmt.Println(c)
}
