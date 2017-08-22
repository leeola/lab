package main

import (
	"github.com/leeola/lab/tui"
	. "github.com/leeola/lab/tui/component"
	. "github.com/leeola/lab/util/component"
)

func main() {
	c := &Box{
		Child: Text("Hello World"),
	}

	if err := tui.Render(AutoQuit(c)); err != nil {
		panic(err)
	}
}
