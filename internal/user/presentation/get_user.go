package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/internal/user/application/find_user"
	"user-service/kit/query"
)

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// GetUserHandler is a handler for getting a user by ID.
// @Summary Get a user by ID
// @Schemes
// @Description Retrieves a user by their UUID
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} GetUserResponse "User found"
// @Failure 400 {object} kit.ErrorResponse "Invalid UUID"
// @Failure 404 {object} kit.ErrorResponse "User not found"
// @Failure 500 {object} kit.ErrorResponse "Internal server error"
// @Router /user/{uuid} [get]
// @Tags user
// @Security ApiKeyAuth
func GetUserHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")

		findUserQuery := find_user.NewFindUserQuery(uuid)

		result, err := queryBus.Ask(findUserQuery)
		if err != nil {
			c.JSON(err.Code, gin.H{"error": err.Error()})
			return
		}

		if result == nil {
			c.JSON(err.Code, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
