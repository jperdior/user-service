package presentation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/internal/user/application"
	"user-service/internal/user/application/update_user"
	"user-service/kit"
)

type updateUserRequest struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

// UpdateUserHandler godoc
// @Summary Update a user
// @Description Update a user's details. Only non-empty fields will be updated.
// @Accept json
// @Produce json
// @Param uuid path string true "User ID"
// @Param updateUserRequest body updateUserRequest true "Update User Request"
// @Success 200 {object} dto.UserDTO "User updated successfully"
// @Failure 400 {object} kit.ErrorResponse
// @Failure 401 {object} kit.ErrorResponse
// @Failure 500 {object} kit.ErrorResponse
// @Router /users/{uuid} [put]
// @Tags user
// @Security Bearer
func UpdateUserHandler(service *update_user.UpdateUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("uuid")
		var request updateUserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authenticatedUser, err := application.GetAuthenticatedUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user, err := service.UpdateUser(
			authenticatedUser,
			uid,
			request.Name,
			request.Email,
			request.Password,
			request.Roles,
		)
		if err != nil {
			var domainError *kit.DomainError
			if errors.As(err, &domainError) {
				c.JSON(domainError.Code, gin.H{"error": domainError.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
