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
