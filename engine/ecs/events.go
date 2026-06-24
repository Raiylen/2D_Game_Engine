package ecs

import ()

type Event struct {
	Name string
}

type Handler func(e Event)

type EventBus struct {
	queue    []Event
	handlers map[string][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		queue:    []Event{},
		handlers: map[string][]Handler{},
	}
}

func (eb *EventBus) Emit(e Event) {
	eb.queue = append(eb.queue, e)
}

func (eb *EventBus) RegisterHandler(eventName string, h Handler) {
	eb.handlers[eventName] = append(eb.handlers[eventName], h)
}

func (eb *EventBus) HandleEvents() {
	for len(eb.queue) > 0 {
		currentEvent := eb.queue[0]
		eb.queue = eb.queue[1:]

		handlers := eb.handlers[currentEvent.Name]
		for _, handler := range handlers {
			handler(currentEvent)
		}
	}
	eb.Reset()
}

func (eb *EventBus) Reset() {
	eb.queue = eb.queue[:0]
}
