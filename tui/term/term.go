package term

// Term is a low level interface for drawing in terminals.
type Term interface {
	Close() error
	Renderer(z uint8) Renderer
}

type Renderer interface {
	Flush() error
	Cell(x, y uint, ru rune) error
}
