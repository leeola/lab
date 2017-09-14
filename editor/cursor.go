package editor

import "github.com/gdamore/tcell"

type Cursor struct {
	screen tcell.Screen
	node   Node
}

func (c Cursor) Options() Options {
	switch c.node.Type {
	case NodeTypeClosure:
		return c.node.Closure.Options
	default:
		return nil
	}
}

func (c Cursor) InsertOption(n int) error {
	drawWord(c.screen, 0, 0,
		c.node.Closure.Options[n].Node.Keyword.Chars)

	return nil
}

func drawWord(s tcell.Screen, x, y int, rs []rune) {
	for i, r := range rs {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}
