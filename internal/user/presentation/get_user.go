package presentation

import (
	"errors"
	"net/http"
	"user-service/internal/user/application/dto"
	"user-service/internal/user/application/find_user"
	"user-service/kit"
	"user-service/kit/query"

	"github.com/gin-gonic/gin"
)

// GetUserHandler is a handler for getting a user by ID.
// @Summary Get a user by ID
// @Schemes
// @Description Retrieves a user by their UUID
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} dto.UserDTO "User found"
// @Failure 400 {object} kit.ErrorResponse "Invalid UUID"
// @Failure 404 {object} kit.ErrorResponse "User not found"
// @Failure 500 {object} kit.ErrorResponse "Internal server error"
// @Router /users/{uuid} [get]
// @Tags user
// @Security Bearer
func GetUserHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("uuid")

		findUserQuery := find_user.NewFindUserQuery(uid)

		user, err := queryBus.Ask(c, findUserQuery)
		if err != nil {
			var domainError *kit.DomainError
			if errors.As(err, &domainError) {
				c.JSON(domainError.Code, gin.H{"error": domainError.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}
		user, ok := user.(*dto.UserDTO)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
