package presentation

import (
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/platform/bus/inmemory"
	"user-service/internal/platform/server/middleware/auth"
	"user-service/internal/user/application/find_users"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit/test/helpers"
	"user-service/kit/test/pages"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsersHandler(t *testing.T) {

	secretKey := "secret"

	repo := new(domainmocks.UserRepository)
	service := find_users.NewFindUsersService(repo)
	queryHandler := find_users.NewFindUsersQueryHandler(service)
	queryBus := inmemory.NewQueryBus()
	queryBus.Register(find_users.FindUsersQueryType, queryHandler)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(auth.JWTMiddleware(secretKey))
	router.GET("/users", GetUsersHandler(queryBus))

	jwtToken, err := helpers.GenerateJwtToken([]string{domain.RoleSuperAdmin}, secretKey)
	require.NoError(t, err)

	userPage := pages.NewUserPage(&jwtToken)

	t.Run("given there's 2 users in the database it returns them", func(t *testing.T) {

		expectedUsers := helpers.CreateManyUsers(2)

		totalRows := int64(len(expectedUsers))

		repo.On("Find", mock.Anything).Return(expectedUsers, totalRows, nil)

		request, err := userPage.GetUsers()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

	})
}
