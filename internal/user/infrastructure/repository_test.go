package infrastructure

import (
	"github.com/google/uuid"
	"testing"
	"user-service/internal/user/domain"
	"user-service/kit/model"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Use SQLite in-memory for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the test database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil
	}
	return db
}

func TestUserRepositoryImpl_FindByID(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)

	user, err := domain.NewUser(
		"peter",
		"test@example.com",
		"password",
	)
	assert.Nil(t, err)

	// Save the user
	err = repo.Save(user)
	assert.Nil(t, err)

	// Retrieve the user by ID
	uuidString, err := uuid.FromBytes(user.ID)
	assert.Nil(t, err)
	foundUser, err := repo.FindByID(uuidString.String())
	assert.Nil(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, "test@example.com", foundUser.Email)
}

func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	id, _ := uuid.New().MarshalBinary()
	// Prepare the test user
	user := &domain.User{
		Base: model.Base{
			ID: id,
		},
		Email: "test@example.com",
	}

	// Save the user
	err := repo.Save(user)
	assert.Nil(t, err)

	// Retrieve the user by email
	foundUser, err := repo.FindByEmail("test@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, "1", foundUser.ID)
}

func TestUserRepositoryImpl_Save(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	id, _ := uuid.New().MarshalBinary()
	// Create a new user
	user := &domain.User{
		Base: model.Base{
			ID: id,
		},
		Email: "newuser@example.com",
	}

	// Save the user
	err := repo.Save(user)
	assert.Nil(t, err)

	// Retrieve and verify user was saved correctly
	foundUser, err := repo.FindByID("2")
	assert.Nil(t, err)
	assert.Equal(t, "newuser@example.com", foundUser.Email)
}
