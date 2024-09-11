package model

import (
	"github.com/google/uuid"
	"time"
	"user-service/kit/event"
)

type Base struct {
	ID        []byte    `gorm:"type:binary(16);primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBase(uid uuid.UUID) (*Base, error) {
	id, err := uid.MarshalBinary() // Generates binary UUID
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	return &Base{
		ID:        id,
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
