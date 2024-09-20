package presentation

import "github.com/gin-gonic/gin"

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

// ForgotPasswordHandler is a handler to recover a user's password
// @Summary Sends a password recovery email to the user
// @Schemes
// @Description Sends a password recovery email to the user
// @Accept  json
// @Produce json
// @Param   login body      login.ForgotPasswordRequest  true  "Login request body"
// @Success 200 {object} map[string]string "Token generated successfully" example({"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."})
// @Failure 400 {object} map[string]string "Invalid input" example({"error": "invalid input"})
// @Failure 401 {object} map[string]string "Invalid credentials" example({"error": "invalid credentials"})
// @Failure 500 {object} map[string]string "Internal server error" example({"error": "internal server error"})
// @Router /forgot-password [post]
// @Tags user
func ForgotPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ForgotPasswordRequest

		// Bind JSON request body to struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		// Call login service with email and password
		// token, err := service.Login(req.Email, req.Password)
		// if err != nil {
		// 	c.JSON(err.Code, gin.H{"error": err.Error()})
		// 	return
		// }

		c.JSON(200, ForgotPasswordResponse{Message: "Password recovery email sent"})
	}
}
