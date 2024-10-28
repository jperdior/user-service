package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user-service/internal/user/application/register"
)

type RegisterUserRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// RegisterUserHandler handles user registration
// @Summary Registers a new user
// @Schemes
// @Description Registers a new user with the provided name, email, and password
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "User registration details" example({ "id": "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0", "name": "Jean Luc Picard", "email": "jlp@federation.com", "password": "enterprise" })
// @Success 201 {object} UserResponse "User registered successfully"
// @Failure 400 {object} kit.ErrorResponse "Invalid input" example {"error": "invalid input"}
// @Failure 500 {object} kit.ErrorResponse "Internal server error" example {"error": "internal server error"}
// @Example request { "id": "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0", "name": "Jean Luc Picard", "email": "jlp@federation.com", "password": "enterprise" }
// @Example 400 { "error": "invalid input" }
// @Example 500 { "error": "internal server error" }
// @Router /register [post]
// @Tags user
// @example request body RegisterUserRequest{ID="6d0f12c8-9fe9-4506-ad59-d386adbbe5c0", Name="Jean Luc Picard", Email="
func RegisterUserHandler(service *register.UserRegisterService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request RegisterUserRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := service.RegisterUser(request.ID, request.Name, request.Email, request.Password)
		if err != nil {
			MapErrorToHTTP(context, err)
			return
		}

		response := UserResponse{
			ID:        request.ID, // Convert binary ID to string if necessary
			Name:      request.Name,
			Email:     request.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		}

		context.JSON(http.StatusCreated, response)
	}
}
