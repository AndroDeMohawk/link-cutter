package event

type Event struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

const (
	LinkVisited = "link.visited"
)

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event, 128),
	}
}

func (e EventBus) Publish(event Event) {
	e.bus <- event

}
func (e EventBus) Subscribe() <-chan Event {
	return e.bus
}
