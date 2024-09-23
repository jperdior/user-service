package presentation

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/user/application/forgot_password"
)

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

// ForgotPasswordHandler handles the password recovery process for a user
// @Summary Send password recovery email
// @Description Sends a password recovery email to the user if the provided email exists in the system
// @Accept  json
// @Produce json
// @Param   forgotPasswordRequest body ForgotPasswordRequest true "Forgot password request body"
// @Success 200 {object} ForgotPasswordResponse "Password recovery email sent"
// @Failure 400 {object} kit.ErrorResponse "Invalid input"
// @Failure 500 {object} kit.ErrorResponse "Internal server error"
// @Router /forgot-password [post]
// @Tags user
func ForgotPasswordHandler(service *forgot_password.ForgotPasswordService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ForgotPasswordRequest

		// Bind JSON request body to struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		err := service.SendResetPasswordEmail(req.Email)
		if err != nil {
			c.JSON(err.Code, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, ForgotPasswordResponse{Message: "Password recovery email sent"})
	}
}
