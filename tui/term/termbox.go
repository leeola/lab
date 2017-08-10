package term

import (
	"sync"

	termbox "github.com/nsf/termbox-go"
)

// New returns a default Term, implemented by termbox.
func New() (Term, error) {
	if err := termbox.Init(); err != nil {
		return nil, err
	}

	return &term{}, nil
}

type term struct {
	flushLock sync.Mutex
}

func (t *term) Close() error {
	termbox.Close()
	return nil
}

func (t *term) Renderer() Renderer {
	return t
}

func (t *term) Flush() error {

	// Termbox seems to have issues with flushing concurrently, so the locking
	// prevents some broken renders and/or panics.
	t.flushLock.Lock()
	defer t.flushLock.Unlock()

	return termbox.Flush()
}

func (t *term) Cell(x, y int, ru rune) error {
	termbox.SetCell(x, y, ru, termbox.ColorDefault, termbox.ColorDefault)
	return nil
}
