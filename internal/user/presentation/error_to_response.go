package presentation

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"user-service/kit"
)

// mapDomainErrorToHTTP maps a DomainError to an HTTP status code and message.
func mapDomainErrorToHTTP(err *kit.DomainError) (int, string) {
	switch err.Key {
	case "user.register.email_exists":
		return http.StatusConflict, err.Message // 409 Conflict
	case "user.register.invalid_email":
		return http.StatusBadRequest, err.Message // 400 Bad Request
	case "user.register.invalid_id":
		return http.StatusBadRequest, err.Message // 400 Bad Request
	case "user.login.invalid_credentials":
		return http.StatusUnauthorized, err.Message // 401 Unauthorized
	case "user.forgot_password.user_not_found":
		return http.StatusNotFound, err.Message // 404 Not Found
	case "user.find_user.error":
		return http.StatusUnauthorized, err.Message // 401 Unauthorized
	default:
		return http.StatusInternalServerError, "Internal server error" // 500 Internal Server Error
	}
}

// MapErrorToHTTP checks the error type, maps it to an HTTP response, and sends the response via the Gin context.
// If the error is a DomainError, it maps using mapDomainErrorToHTTP; otherwise, it defaults to a 500 status code.
func MapErrorToHTTP(c *gin.Context, err error) {
	var domainErr *kit.DomainError
	if errors.As(err, &domainErr) {
		statusCode, message := mapDomainErrorToHTTP(domainErr)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}
	// For non-domain errors, return a 500 Internal Server Error.
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}
