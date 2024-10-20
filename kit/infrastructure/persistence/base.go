package persistence

import (
	"time"
	"user-service/kit/domain"
)

type Base struct {
	ID        []byte    `gorm:"type:binary(16);primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBase(id domain.UuidValueObject) (*Base, error) {
	currentTime := time.Now()
	return &Base{
		ID:        id.Bytes(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}, nil
}
