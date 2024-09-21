package kit

import (
	"fmt"
	"log"
)

type DomainError struct {
	Message string
	Key     string
	Code    int
}

// Error implements the error interface for DomainError.
func (e *DomainError) Error() string {
	return fmt.Sprintf("DomainError{Message: %s, Key: %s, Code: %d}", e.Message, e.Key, e.Code)
}

func NewDomainError(message, key string, code int) *DomainError {
	domainErr := &DomainError{
		Message: message,
		Key:     key,
		Code:    code,
	}
	log.Println(domainErr) // Log the error
	return domainErr
}

type ErrorResponse struct {
	Error string `json:"error"`
}
