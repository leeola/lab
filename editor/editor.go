package editor

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/leeola/lab/editor/mode"
)

type Editor struct {
	screen tcell.Screen
	quitC  chan struct{}
	cursor *Cursor
	mode   mode.Mode
}

func New(n Node) (*Editor, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err = s.Init(); err != nil {
		return nil, err
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Clear()

	return &Editor{
		screen: s,
		quitC:  make(chan struct{}),
		mode:   mode.Input,
		cursor: &Cursor{
			screen: s,
			node:   n,
		},
	}, nil
}

func (e *Editor) eventLoop() {
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

func (e *Editor) renderLoop() {
	for {
		select {
		case <-e.quitC:
			return
		case <-time.After(time.Millisecond * 50):
		}

		e.screen.Show()
	}
}

func (e *Editor) Start() error {
	go e.eventLoop()
	e.renderLoop()
	e.screen.Fini()
	return nil
}
