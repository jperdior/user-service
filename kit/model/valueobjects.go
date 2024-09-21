package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

// EmailValueObject represents a value object for emails
type EmailValueObject string

func NewEmailValueObject(value string) (EmailValueObject, error) {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(value) {
		return "", errors.New("invalid email format")
	}
	return EmailValueObject(value), nil
}

func (emailValueObject *EmailValueObject) Value() string {
	return string(*emailValueObject)
}

// UuidValueObject represents a value object for UUIDs
type UuidValueObject uuid.UUID

func NewUuidValueObject(value string) (UuidValueObject, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return UuidValueObject{}, err
	}
	return UuidValueObject(uid), nil
}

func (uidValueObject *UuidValueObject) Scan(value interface{}) error {
	if value == nil {
		*uidValueObject = UuidValueObject{} // If value is nil, set it to the zero value
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan UUID: expected []byte but got %T", value)
	}

	parsedUUID, err := uuid.FromBytes(bytes)
	if err != nil {
		return fmt.Errorf("failed to parse UUID from bytes: %w", err)
	}

	*uidValueObject = UuidValueObject(parsedUUID)
	return nil
}

func (uidValueObject *UuidValueObject) Value() (driver.Value, error) {
	return uuid.UUID(*uidValueObject).MarshalBinary()
}

func (uidValueObject *UuidValueObject) String() string {
	return uuid.UUID(*uidValueObject).String()
}

func (uidValueObject *UuidValueObject) Bytes() []byte {
	bytes, _ := uuid.UUID(*uidValueObject).MarshalBinary()
	return bytes
}
