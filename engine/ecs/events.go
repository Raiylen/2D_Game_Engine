package ecs

import ()

type Event struct {
	Name string
	Data any
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
	for _, event := range eb.queue {
		for _, h := range eb.handlers[event.Name] {
			h(event)
		}
	}
	eb.queue = eb.queue[:0]
}
