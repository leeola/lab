package editor

import (
	"github.com/gdamore/tcell"
	"github.com/leeola/lab/editor/event"
)

func (e *Editor) PollEvent(e *event.Event) {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				close(e.quitC)
				return
			case tcell.KeyRune:
				e.HandleInput(ev.Rune())
			case tcell.KeyCtrlL:
				e.screen.Sync()
			}
		case *tcell.EventResize:
			e.screen.Sync()
		}
	}
}

func (e *Editor) HandleMode(key rune) {
}
