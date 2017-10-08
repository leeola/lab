package mode

import "fmt"

type Mode int

const (
	Unknown Mode = iota
	NodeNavigation
	Input
)

func (m Mode) String() string {
	switch m {
	case Unknown:
		return "Unknown"
	case NodeNavigation:
		return "NodeNavigation"
	case Input:
		return "Input"
	default:
		return fmt.Sprintf("invalid mode: %d", int(m))
	}
}
