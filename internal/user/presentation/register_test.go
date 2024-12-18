package presentation

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/user/application/register"
	"user-service/internal/user/domain"
	"user-service/internal/user/domain/domainmocks"
	"user-service/kit/event/eventmocks"
	"user-service/kit/test/pages"
)

func TestRegisterUserHandler(t *testing.T) {
	repo := new(domainmocks.UserRepository)
	eventBus := new(eventmocks.Bus)
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	userService := register.NewUserRegisterService(repo, eventBus)
	router.POST("/register", RegisterUserHandler(userService))

	userPage := pages.NewUserPage(nil)

	t.Run("given a valid request it returns 201", func(t *testing.T) {

		repo.On("FindByEmail", "jlp@federation.com").Return(nil, nil)
		repo.On("Save", mock.Anything).Return(nil)
		eventBus.On("Publish", mock.Anything).Return(nil)

		payload := map[string]string{
			"id":       "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0",
			"name":     "Jean Luc Picard",
			"email":    "jlp@federation.com",
			"password": "enterprise",
		}

		req, err := userPage.RegisterUser(payload)

		// Create a ResponseRecorder to record the response
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Check the response
		assert.Equal(t, http.StatusCreated, recorder.Code)

		// Check the response body
		var response UserResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, payload["id"], response.ID)
		assert.Equal(t, payload["name"], response.Name)
		assert.Equal(t, payload["email"], response.Email)
	})

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		// Define the request payload
		payload := map[string]string{
			"id":       "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0",
			"name":     "Jean Luc Picard",
			"email":    "federation.com",
			"password": "enterprise",
		}

		req, err := userPage.RegisterUser(payload)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		// Check the response body
		var response map[string]string
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"], "Invalid email format")
	})

	t.Run("given an existing user it returns 409", func(t *testing.T) {

		repo.On("FindByEmail", "first@federation.com").Return(&domain.User{}, nil)

		// Define the request payload
		payload := map[string]string{
			"id":       "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0",
			"name":     "William Riker",
			"email":    "first@federation.com",
			"password": "enterprise",
		}

		req, err := userPage.RegisterUser(payload)
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusConflict, recorder.Code)
		// Check the response body
		var response map[string]string
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"], "User with email already exists")
	})

}
