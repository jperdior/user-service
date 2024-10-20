package presentation

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/user/application/forgot_password"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit/test/helpers"
)

func TestForgotPasswordHandler(t *testing.T) {

	emailService := new(domainmocks.EmailService)
	userRepositoryMock := new(domainmocks.UserRepository)
	tokenMock := new(domainmocks.TokenService)
	forgotPasswordService := forgot_password.NewForgotPasswordService(userRepositoryMock, emailService, tokenMock)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/forgot-password", ForgotPasswordHandler(forgotPasswordService))

	t.Run("Test correct email", func(t *testing.T) {
		expectedUser := helpers.CreateUser()
		resetPasswordToken := "resetPasswordToken"
		userRepositoryMock.On("FindByEmail", expectedUser.Email.Value()).Return(expectedUser, nil)
		tokenMock.On("GenerateResetPasswordToken", expectedUser).Return(resetPasswordToken, nil)
		emailService.On("SendPasswordResetEmail", expectedUser.Email.Value(), resetPasswordToken).Return(nil)

		payload := map[string]string{
			"email": expectedUser.Email.Value(),
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/forgot-password", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		userRepositoryMock.AssertExpectations(t)
		emailService.AssertExpectations(t)

	})

	t.Run("Test incorrect email", func(t *testing.T) {

		userRepositoryMock.On("FindByEmail", "wr@federation.com").Return(nil, errors.New("user not found"))

		payload := map[string]string{
			"email": "wr@federation.com",
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/forgot-password", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
		userRepositoryMock.AssertExpectations(t)
	})

}
