package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-service/internal/user/application/login"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
)

func TestLoginUserHandler(t *testing.T) {

	repo := new(domainmocks.UserRepository)
	tokenService := new(domainmocks.TokenService)
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	userService := login.NewUserLoginService(repo, tokenService)
	router.POST("/login", LoginUserHandler(userService))

	t.Run("Test correct email and password", func(t *testing.T) {

		repo.On("FindByEmail", "jlp@federation.com").Return(&domain.User{}, nil)
		tokenService.On("GenerateToken", mock.Anything).Return("token", nil)

		payload := "email=jlp@federation.com&password=enterprise"
		req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		require.Equal(t, http.StatusOK, recorder.Code)
		tokenService.AssertExpectations(t)
	})

}
