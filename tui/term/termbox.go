package term

import (
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
}

func (t *term) Close() error {
	termbox.Close()
	return nil
}

func (t *term) Renderer() Renderer {
	return t
}

func (t *term) Flush() error {
	return termbox.Flush()
}

func (t *term) Cell(x, y int, ru rune) error {
	termbox.SetCell(x, y, ru, termbox.ColorDefault, termbox.ColorDefault)
	return nil
}
