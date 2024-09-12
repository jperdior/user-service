package presentation

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/user/application/login"
)

// LoginUserHandler is a handler to login a user
// @Summary Authenticates a user with the provided email and password
// @Schemes
// @Description Authenticates a user and returns a JWT token if credentials are valid
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "User email" example(jlp@federation.com)
// @Param password formData string true "User password" example(enterprise)
// @Success 200 {object} map[string]string "Token generated successfully" example({"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."})
// @Failure 400 {object} map[string]string "Invalid input" example({"error": "invalid input"})
// @Failure 401 {object} map[string]string "Invalid credentials" example({"error": "invalid credentials"})
// @Failure 500 {object} map[string]string "Internal server error" example({"error": "internal server error"})
// @Router /login [post]
// @Tags user
func LoginUserHandler(service *login.UserLoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		token, err := service.Login(email, password)
		if err != nil {
			c.JSON(err.Code, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}
