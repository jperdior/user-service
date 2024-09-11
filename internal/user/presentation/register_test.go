package presentation

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/user/application"
	"user-service/kit/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"user-service/internal/user/domain"
	"user-service/internal/user/infrastructure"
)

func setupTestRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	userRepo := infrastructure.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	router.POST("/register", RegisterUserHandler(userService))

	return router
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil
	}
	return db
}

func TestRegisterUserHandler(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	insertUserFixture := func(email string) {
		id, _ := uuid.New().MarshalBinary()
		// Prepare the test user
		user := &domain.User{
			Base: model.Base{
				ID: id,
			},
			Email:    email,
			Name:     "Test User",
			Password: "password",
			Roles:    domain.UserRoles{domain.RoleUser},
		}
		if err := db.Create(&user).Error; err != nil {
			t.Fatalf("Failed to insert user fixture: %v", err)
		}
	}

	t.Run("given a valid request it returns 201", func(t *testing.T) {

		// Define the request payload
		payload := map[string]string{
			"id":       "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0",
			"name":     "Jean Luc Picard",
			"email":    "jlp@federation.com",
			"password": "enterprise",
		}

		body, _ := json.Marshal(payload)

		// Create a request to send to the handler
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder to record the response
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Check the response
		assert.Equal(t, http.StatusCreated, recorder.Code)

		// Check the response body
		var response UserResponse
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, payload["id"], response.ID)
		assert.Equal(t, payload["name"], response.Name)
		assert.Equal(t, payload["email"], response.Email)

		// Check if the user is in the database
		var user domain.User
		err = db.First(&user, "email = ?", payload["email"]).Error
		require.NoError(t, err)
		assert.Equal(t, payload["name"], user.Name)
		assert.Equal(t, payload["email"], user.Email)
	})

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		// Define the request payload
		payload := map[string]string{
			"id":       "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0",
			"name":     "Jean Luc Picard",
			"email":    "federation.com",
			"password": "enterprise",
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		// Check the response body
		var response map[string]string
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"], "Invalid email format")
	})

	t.Run("given an existing user it returns 400", func(t *testing.T) {
		// arrange
		insertUserFixture("first@federation.com")

		// Define the request payload
		payload := map[string]string{
			"id":       "6d0f12c8-9fe9-4506-ad59-d386adbbe5c0",
			"name":     "William Riker",
			"email":    "first@federation.com",
			"password": "enterprise",
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		// Check the response body
		var response map[string]string
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Contains(t, response["error"], "User with email already exists")
	})

}
