package persistence

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"user-service/internal/user/domain"
	kitDomain "user-service/kit/domain"
)

func setupTestDB() *gorm.DB {
	// Use SQLite in-memory for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the test database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&UserModel{})
	if err != nil {
		return nil
	}
	return db
}

func TestUserRepositoryImpl_FindByID(t *testing.T) {
	db := setupTestDB()
	repo := NewGormUserRepository(db)

	uid := kitDomain.RandomUuidValueObject()
	// Prepare the test user
	user := &domain.User{
		ID:    uid,
		Email: "test@example.com",
	}

	// Save the user
	err := repo.Save(user)
	require.NoError(t, err)

	// Retrieve the user by ID
	require.NoError(t, err)
	uidString := uid.String()
	foundUser, err := repo.FindByID(uidString)
	require.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, "test@example.com", foundUser.Email.Value())
}

func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	db := setupTestDB()
	repo := NewGormUserRepository(db)
	uid := kitDomain.RandomUuidValueObject()
	// Prepare the test user
	user := &domain.User{
		ID:    uid,
		Email: "test@example.com",
	}

	// Save the user
	err := repo.Save(user)
	require.NoError(t, err)

	// Retrieve the user by email
	foundUser, err := repo.FindByEmail("test@example.com")
	require.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, uid.String(), foundUser.ID.String())
}

func TestUserRepositoryImpl_Save(t *testing.T) {
	db := setupTestDB()
	repo := NewGormUserRepository(db)
	uid := kitDomain.RandomUuidValueObject()
	// Prepare the test user
	user := &domain.User{
		ID:    uid,
		Email: "newuser@example.com",
	}

	// Save the user
	err := repo.Save(user)
	require.NoError(t, err)

	// Retrieve and verify user was saved correctly
	require.NoError(t, err)
	uidString := uid.String()
	foundUser, err := repo.FindByID(uidString)
	require.NoError(t, err)
	assert.Equal(t, "newuser@example.com", foundUser.Email.Value())
}
