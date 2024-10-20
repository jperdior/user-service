package domain

import (
	"user-service/kit/event"
)

type BaseAggregate struct {
	events []event.Event
}

func (a *BaseAggregate) PullEvents() []event.Event {
	events := a.events
	a.events = []event.Event{}
	return events
}

func (a *BaseAggregate) Record(event event.Event) {
	a.events = append(a.events, event)
}
