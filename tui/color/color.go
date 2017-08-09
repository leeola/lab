package color

type Attribute uint16

type Color Attribute

const (
	Default Color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	AttrBold Attribute = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)
