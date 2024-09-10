package model

import "user-service/kit/event"

type Base struct {
	ID        []byte `gorm:"type:binary(16);primary_key"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BaseAggregate struct {
	events []event.Event `gorm:"-"`
}

func (a *BaseAggregate) PullEvents() []event.Event {
	events := a.events
	a.events = []event.Event{}
	return events
}
