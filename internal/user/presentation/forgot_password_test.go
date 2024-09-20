package presentation

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/user/application/forgot_password"
	"user-service/internal/user/domain/domainmocks"
)

func TestForgotPasswordHandler(t *testing.T) {

	emailServiceMock := new(domainmocks.EmailService)
	userRepositoryMock := new(domainmocks.UserRepository)
	forgotPasswordService := forgot_password.NewForgotPasswordService(userRepositoryMock, emailServiceMock)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/forgot-password", ForgotPasswordHandler(forgotPasswordService))

	t.Run("Test correct email", func(t *testing.T) {
		userRepositoryMock.On("FindByEmail", "jlp@federation.com").Return(nil, nil)
		emailServiceMock.On("SendPasswordResetEmail", "jlp@federation.com").Return(nil)

		payload := map[string]string{
			"email": "jlp@federation.com",
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/forgot-password", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		userRepositoryMock.AssertExpectations(t)
		emailServiceMock.AssertExpectations(t)

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
		emailServiceMock.AssertNotCalled(t, "SendPasswordResetEmail", mock.Anything)
	})

}
