package presentation

import (
	"net/http"
	"user-service/internal/user/application/find_user"
	"user-service/internal/user/domain"
	"user-service/kit/query"

	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
		uid := c.Param("uuid")

		findUserQuery := find_user.NewFindUserQuery(uid)

		user, err := queryBus.Ask(findUserQuery)
		if err != nil {
			c.JSON(err.Code, gin.H{"error": err.Error()})
			return
		}
		userEntity, ok := user.(*domain.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
			return
		}

		userResponse := GetUserResponse{
			ID:        userEntity.GetID(),
			Email:     userEntity.Email,
			CreatedAt: userEntity.CreatedAt.String(),
			UpdatedAt: userEntity.UpdatedAt.String(),
		}
		c.JSON(http.StatusOK, userResponse)
	}
}
