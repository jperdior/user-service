package inmemory

import (
	"log"
	"user-service/kit/event"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	handlers map[event.Type][]event.Handler
}

// NewEventBus initializes a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

// Publish implements the event.Bus interface.
func (b *EventBus) Publish(events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler := handler
			go func() {
				err := handler.Handle(evt)
				if err != nil {
					log.Printf("Error while handling %s - %s\n", evt.Type(), err)
				}
			}()
		}
	}

	return nil
}
