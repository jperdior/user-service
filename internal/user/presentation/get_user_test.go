package presentation

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/platform/bus/inmemory"
	"user-service/internal/platform/server/middleware/auth"
	"user-service/internal/platform/token"
	"user-service/internal/user/application/find_user"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit/model"

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
	router.GET("/user/:uuid", GetUserHandler(queryBus)).Use(auth.JWTMiddleware(secretKey))

	testerUseriD := "7d8a8225-73da-4cc2-97fd-70d8e3baf6ac"
	testerUid, err := model.NewUuidValueObject(testerUseriD)
	require.NoError(t, err)

	baseUser := model.Base{
		ID: testerUid.Bytes(),
	}
	authenticatedUser := domain.User{
		Base:  baseUser,
		Email: "tester@federation.com",
	}

	tokenService := token.NewJwtService(secretKey, 1)
	token, err := tokenService.GenerateToken(&authenticatedUser)
	require.NoError(t, err)

	t.Run("given a user id it should return the user", func(t *testing.T) {

		userID := "b167da12-7bc7-4234-99d2-5d4e43886975"
		uid, err := model.NewUuidValueObject(userID)
		require.NoError(t, err)

		baseUser := model.Base{
			ID: uid.Bytes(),
		}
		expectedUser := domain.User{
			Base:  baseUser,
			Email: "jlp@federation.com",
		}

		repo.On("FindByID", uid.String()).Return(&expectedUser, nil)

		request, err := http.NewRequest(http.MethodGet, "/user/"+userID, nil)
		require.NoError(t, err)
		request.Header.Set("Authorization", "Bearer "+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
		responseBody := recorder.Body.String()
		assert.Contains(t, responseBody, userID)
		assert.Contains(t, responseBody, expectedUser.Email)

	})

	t.Run("given a user id that does not exist it should return a not found error", func(t *testing.T) {

		userID := "1e10c93e-eb59-4562-9dc6-621157c7458e"

		repo.On("FindByID", userID).Return(nil, nil)

		request, err := http.NewRequest(http.MethodGet, "/user/"+userID, nil)
		require.NoError(t, err)
		request.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusNotFound, recorder.Code)

	})

	t.Run("given an invalid user id it should return a bad request error", func(t *testing.T) {
		userID := "FDASFSDF"
		request, err := http.NewRequest(http.MethodGet, "/user/"+userID, nil)
		require.NoError(t, err)
		request.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})

}
