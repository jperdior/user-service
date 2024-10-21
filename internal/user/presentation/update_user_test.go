package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/platform/server/middleware/auth"
	"user-service/internal/user/application/update_user"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit/test/helpers"
	"user-service/kit/test/pages"
)

func TestUpdateUserHandler(t *testing.T) {

	const secretKey = "secret"

	repo := new(domainmocks.UserRepository)
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	service := update_user.NewUpdateUserService(repo)
	router.Use(auth.JWTMiddleware(secretKey))
	router.PUT("/users/:uuid", UpdateUserHandler(service))

	type testCase struct {
		name         string
		roles        []string
		expectedUser *domain.User
		payload      map[string]string
		expected     int
	}

	testCases := []testCase{
		{
			name:         "Update user as super admin",
			roles:        []string{domain.RoleSuperAdmin},
			expectedUser: helpers.CreateUser(),
			payload: map[string]string{
				"name":  "Jean-Luc Picard",
				"email": "jlp@romulans.com",
			},
			expected: http.StatusOK,
		},
		{
			name:         "Update user as regular user",
			roles:        []string{domain.RoleUser},
			expectedUser: helpers.CreateUser(),
			payload: map[string]string{
				"name":  "Jean-Luc Picard",
				"email": "jlp@romulans.com",
			},
			expected: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			jwtToken, err := helpers.GenerateJwtToken(tc.roles, secretKey)
			require.NoError(t, err)

			userPage := pages.NewUserPage(&jwtToken)

			repo.On("FindByID", tc.expectedUser.ID).Return(tc.expectedUser, nil)
			repo.On("Save", tc.expectedUser).Return(nil)

			req, err := userPage.UpdateUser(tc.expectedUser.GetID(), tc.payload)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			require.Equal(t, tc.expected, recorder.Code)

		})
	}
}
