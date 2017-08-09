package term

// Term is a low level interface for drawing in terminals.
type Term interface {
	Close() error
	Renderer() Renderer
}

type Renderer interface {
	Flush() error
	Cell(x, y int, ru rune) error
}
