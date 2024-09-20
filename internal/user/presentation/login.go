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
// @Param   login body      login.LoginRequest  true  "Login request body"
// @Success 200 {object} map[string]string "Token generated successfully" example({"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."})
// @Failure 400 {object} map[string]string "Invalid input" example({"error": "invalid input"})
// @Failure 401 {object} map[string]string "Invalid credentials" example({"error": "invalid credentials"})
// @Failure 500 {object} map[string]string "Internal server error" example({"error": "internal server error"})
// @Router /login [post]
// @Tags user
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
