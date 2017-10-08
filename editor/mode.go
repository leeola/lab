package editor

import "github.com/leeola/lab/editor/mode"

// HandleInput provides programmatic input access to the editor.
//
// TODO(leeola): include modifiers in this arg.
// Likely it should be a editor/key.Key type.
func (e *Editor) HandleInput(key rune) {
	switch e.mode {
	case mode.NodeNavigation:
	case mode.Input:
		e.cursor.InsertOption(0)
	}
}

func (e *Editor) SetMode(m mode.Mode) {
	e.mode = m

	var mName []rune
	switch m {
	case mode.NodeNavigation:
		mName = []rune("navigation")
	case mode.Input:
		mName = []rune("input")
	default:
		mName = []rune(m.String())
	}

	drawWord(e.screen, e.width-len(mName), e.height-1, mName)
}
