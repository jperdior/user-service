package model

import (
	"time"
	"user-service/kit/event"
)

type Base struct {
	ID        []byte    `gorm:"type:binary(16);primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBase(id UuidValueObject) (*Base, error) {
	currentTime := time.Now()
	return &Base{
		ID:        id.Bytes(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}, nil
}

type BaseAggregate struct {
	events []event.Event `gorm:"-"`
}

func (a *BaseAggregate) PullEvents() []event.Event {
	events := a.events
	a.events = []event.Event{}
	return events
}

func (a *BaseAggregate) Record(event event.Event) {
	a.events = append(a.events, event)
}
