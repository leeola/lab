package event

type Type uint8

const (
	TypeUnknown Type = iota
	TypeKey
	TypeRune
)

type Key uint16

const (
	KeyUnknown Key = iota
	KeyEnter
)

type Event struct {
	Type Type
	Key  Key
	Rune rune
}
