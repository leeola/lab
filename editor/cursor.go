package editor

import "github.com/gdamore/tcell"

type Cursor struct {
	screen tcell.Screen
	node   Node
	x, y   int
}

func (c Cursor) Options() Options {
	switch c.node.Type {
	case NodeTypeClosure:
		return c.node.Closure.Options
	default:
		return nil
	}
}

func (c *Cursor) InsertOption(n int) error {
	c.x = drawNode(c.screen, c.x, c.y, c.node.Closure.Options[n].Node)

	return nil
}

func drawNode(s tcell.Screen, x, y int, n Node) (after_x int) {
	switch n.Type {
	case NodeTypeGroup:
		for _, n := range n.Group.Nodes {
			x = drawNode(s, x, y, n)
		}
		return x

	case NodeTypeKeyword:
		drawWord(s, x, y, n.Keyword.Chars)
		return x + int(n.Keyword.Width)

	default:
		return x
	}
}

func drawWord(s tcell.Screen, x, y int, rs []rune) {
	for i, r := range rs {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}
