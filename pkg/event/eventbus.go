package event

type EventType int

const (
	EventLinkVisited = iota
)

type Event struct {
	Type EventType
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{bus: make(chan Event)}
}

func (e *EventBus) Publish(event Event) {
	e.bus <- event
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.bus
}
