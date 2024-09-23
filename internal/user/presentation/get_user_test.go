package presentation

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/platform/bus/inmemory"
	"user-service/internal/platform/server/middleware/auth"
	"user-service/internal/user/application/find_user"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit/test/helpers"
	"user-service/kit/test/pages"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserHandler(t *testing.T) {

	secretKey := "secret"

	repo := new(domainmocks.UserRepository)
	service := find_user.NewUserFinderService(repo)
	queryHandler := find_user.NewFindUserQueryHandler(service)
	queryBus := inmemory.NewQueryBus()
	queryBus.Register(find_user.FindUserQueryType, queryHandler)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(auth.JWTMiddleware(secretKey))
	router.GET("/users/:uuid", GetUserHandler(queryBus))

	jwtToken, err := helpers.GenerateJwtToken([]string{domain.RoleSuperAdmin}, secretKey)
	require.NoError(t, err)

	userPage := pages.NewUserPage(&jwtToken)

	t.Run("given a user id it should return the user", func(t *testing.T) {

		expectedUser := helpers.CreateUser()

		repo.On("FindByID", expectedUser.GetID()).Return(expectedUser, nil)

		request, err := userPage.GetUser(expectedUser.GetID())
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
		responseBody := recorder.Body.String()
		assert.Contains(t, responseBody, expectedUser.GetID())
		assert.Contains(t, responseBody, expectedUser.Email)

	})

	t.Run("given a user id that does not exist it should return a not found error", func(t *testing.T) {

		userID := "1e10c93e-eb59-4562-9dc6-621157c7458e"

		repo.On("FindByID", userID).Return(nil, nil)

		request, err := userPage.GetUser(userID)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusNotFound, recorder.Code)

	})

	t.Run("given an invalid user id it should return a bad request error", func(t *testing.T) {
		userID := "FDASFSDF"
		request, err := userPage.GetUser(userID)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})

}
