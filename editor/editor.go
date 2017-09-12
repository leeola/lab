package editor

import (
	"time"

	"github.com/gdamore/tcell"
)

type Editor struct {
	screen tcell.Screen
	quitC  chan struct{}
}

func New() (*Editor, error) {
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
		case <-time.After(time.Second):
		}
	}
}

func (e *Editor) Start() error {
	go e.eventLoop()
	e.renderLoop()
	e.screen.Fini()
	return nil
}
