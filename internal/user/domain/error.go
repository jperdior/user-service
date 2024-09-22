package domain

import (
	"net/http"
	"user-service/kit"
)

// NewExistingUserError creates a new ExistingUserError.
func NewExistingUserError() *kit.DomainError {
	return kit.NewDomainError(
		"User with email already exists",
		"user.register.email_exists",
		http.StatusBadRequest)
}

// NewInvalidEmailError creates a new InvalidEmailError.
func NewInvalidEmailError() *kit.DomainError {
	return kit.NewDomainError(
		"Invalid email format",
		"user.register.invalid_email",
		http.StatusBadRequest)
}

// NewInvalidIDError creates a new InvalidIDError.
func NewInvalidIDError() *kit.DomainError {
	return kit.NewDomainError(
		"Invalid ID format",
		"user.register.invalid_id",
		http.StatusBadRequest)
}

// NewInvalidCredentialsError creates a new InvalidCredentialsError.
func NewInvalidCredentialsError() *kit.DomainError {
	return kit.NewDomainError(
		"Invalid credentials",
		"user.login.invalid_credentials",
		http.StatusUnauthorized)
}

// NewUserNotFoundError creates a new UserNotFoundError.
func NewUserNotFoundError() *kit.DomainError {
	return kit.NewDomainError(
		"User not found",
		"user.forgot_password.user_not_found",
		http.StatusNotFound)
}

// NewUnauthorizedError creates a new UnauthorizedError.
func NewUnauthorizedError() *kit.DomainError {
	return kit.NewDomainError(
		"Unauthorized",
		"user.find_user.error",
		http.StatusUnauthorized)
}
