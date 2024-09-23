package presentation

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/user/application/login"
)

// LoginRequest represents the expected login payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response payload for a successful login
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginUserHandler is a handler to login a user
// @Summary Authenticates a user with the provided email and password
// @Schemes
// @Description Authenticates a user and returns a JWT token if credentials are valid
// @Accept  json
// @Produce json
// @Param   login body      LoginRequest true "User login details"
// @Success 200 {object} 	LoginResponse
// @Failure 400 {object} kit.ErrorResponse "Invalid input"
// @Failure 401 {object} kit.ErrorResponse "Invalid credentials"
// @Failure 500 {object} kit.ErrorResponse "Internal server error"
// @Router /login [post]
// @Tags user
// @example login body LoginRequest{Email="julio.perdiguer@gmail.com", Password="test"}
func LoginUserHandler(service *login.UserLoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest

		// Bind JSON request body to struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		// Call login service with email and password
		token, err := service.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(err.Code, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, LoginResponse{Token: token})
	}
}
