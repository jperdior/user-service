package kit

import (
	"errors"
	"github.com/google/uuid"
)

func UuidStringToBinary(id string) ([]byte, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}
	return parsedUUID.MarshalBinary()
}
