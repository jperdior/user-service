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
	"user-service/internal/user/application/login"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit"
)

func TestLoginUserHandler(t *testing.T) {

	repo := new(domainmocks.UserRepository)
	tokenService := new(domainmocks.TokenService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	userService := login.NewUserLoginService(repo, tokenService)
	router.POST("/login", LoginUserHandler(userService))

	uid, err := kit.NewUuidValueObject("6d0f12c8-9fe9-4506-ad59-d386adbbe5c0")
	require.NoError(t, err)
	email, err := kit.NewEmailValueObject("jlp@federation.com")

	t.Run("Test correct email and password", func(t *testing.T) {

		mockUser, err := domain.NewUser(
			uid,
			"Jean Luc Picard",
			email,
			"enterprise",
		)

		repo.On("FindByEmail", "jlp@federation.com").Return(mockUser, nil)
		tokenService.On("GenerateToken", mock.Anything).Return("token", nil)

		payload := map[string]string{
			"email":    "jlp@federation.com",
			"password": "enterprise",
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)
		tokenService.AssertExpectations(t)

		var response LoginResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "token", response.Token)
	})

	t.Run("Test incorrect email", func(t *testing.T) {

		repo.On("FindByEmail", "wr@federation.com").Return(nil, errors.New("some error"))

		payload := map[string]string{
			"email":    "wr@federation.com",
			"password": "enterprise",
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("Test incorrect password", func(t *testing.T) {
		mockUser, err := domain.NewUser(
			uid,
			"Jean Luc Picard",
			email,
			"enterprise",
		)

		repo.On("FindByEmail", "jlp@federation.com").Return(mockUser, nil)

		payload := map[string]string{
			"email":    "jlp@federation.com",
			"password": "ferengi",
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
}
