package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/user/application/find_user"
	"user-service/internal/user/domain"
	"user-service/kit"
	"user-service/kit/model"
	"user-service/kit/query/querymocks"
)

func TestGetUserHandler(t *testing.T) {

	queryBus := new(querymocks.Bus)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/user/:uuid", GetUserHandler(queryBus))

	t.Run("given a user id it should return the user", func(t *testing.T) {

		userID := "b167da12-7bc7-4234-99d2-5d4e43886975"
		uid, err := kit.NewUuidValueObject(userID)
		require.NoError(t, err)
		baseUser := model.Base{
			ID: uid.Bytes(),
		}
		expectedUser := domain.User{
			Base:  baseUser,
			Email: "jlp@federation.com",
		}
		queryBus.On("Ask", find_user.NewFindUserQuery(userID)).Return(expectedUser, nil)

		request, _ := http.NewRequest(http.MethodGet, "/user/"+userID, nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
		queryBus.AssertExpectations(t)
	})

	t.Run("given a user id that does not exist it should return a not found error", func(t *testing.T) {

		userID := "b167da12-7bc7-4234-99d2-5d4e43886975"
		queryBus.On("Ask", find_user.NewFindUserQuery(userID)).Return(nil, domain.NewUserNotFoundError())

		request, _ := http.NewRequest(http.MethodGet, "/user/"+userID, nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
		queryBus.AssertExpectations(t)
	})

	t.Run("given an invalid user id it should return a bad request error", func(t *testing.T) {

		userID := "invalid-uuid"
		request, _ := http.NewRequest(http.MethodGet, "/user/"+userID, nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		queryBus.AssertExpectations(t)
	})

}
