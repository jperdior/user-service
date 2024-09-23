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
	"user-service/kit/model"
	"user-service/kit/test/pages"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"user-service/internal/platform/token"
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

	testerUserId := "7d8a8225-73da-4cc2-97fd-70d8e3baf6ac"
	testerUid, err := model.NewUuidValueObject(testerUserId)
	require.NoError(t, err)

	baseUser := model.Base{
		ID: testerUid.Bytes(),
	}
	authenticatedUser := domain.User{
		Base:  baseUser,
		Name:  "Tester",
		Email: "tester@federation.com",
		Roles: []string{domain.RoleSuperAdmin},
	}

	tokenService := token.NewJwtService(secretKey, 1)
	jwtToken, err := tokenService.GenerateToken(&authenticatedUser)
	require.NoError(t, err)

	userPage := pages.NewUserPage(&jwtToken)

	t.Run("given there's 2 users in the database it returns them", func(t *testing.T) {

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

		user2ID := "b167da12-7bc7-4234-99d2-5d4e43886976"
		uid2, err := model.NewUuidValueObject(user2ID)
		require.NoError(t, err)

		baseUser2 := model.Base{
			ID: uid2.Bytes(),
		}
		expectedUser2 := domain.User{
			Base:  baseUser2,
			Email: "wr@federation.com",
		}

		totalRows := int64(2)

		repo.On("Find", mock.Anything).Return([]*domain.User{&expectedUser, &expectedUser2}, totalRows, nil)

		request, err := userPage.GetUsers()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

	})
}
