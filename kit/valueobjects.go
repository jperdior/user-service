package kit

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
)

// EmailValueObject represents a value object for emails
type EmailValueObject struct {
	value string
}

func NewEmailValueObject(value string) (EmailValueObject, error) {
	err := validateValue(value)
	if err != nil {
		return EmailValueObject{}, err
	}
	return EmailValueObject{value: value}, nil
}

func validateValue(value string) error {
	// Regular expression for basic email validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(value) {
		return errors.New("invalid email format")
	}
	return nil
}

func (e EmailValueObject) Value() string {
	return e.value
}

// UuidValueObject represents a value object for UUIDs
type UuidValueObject struct {
	value uuid.UUID
}

func NewUuidValueObject(value string) (UuidValueObject, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return UuidValueObject{}, err
	}
	return UuidValueObject{value: uid}, nil
}

func (u UuidValueObject) Value() uuid.UUID {
	return u.value
}
