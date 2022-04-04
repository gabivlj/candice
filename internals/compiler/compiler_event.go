package compiler

type EventType = int

const (
	AddFlags EventType = iota
)

type Event struct {
	Kind EventType
	Data string
}
