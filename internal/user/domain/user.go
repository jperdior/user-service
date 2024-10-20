package domain

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user-service/kit/domain"
)

const (
	RoleUser       = "ROLE_USER"
	RoleSuperAdmin = "ROLE_SUPER_ADMIN"
)

type UserRoles []string

type User struct {
	domain.BaseAggregate
	ID            domain.UuidValueObject
	Name          string
	Email         domain.EmailValueObject
	Password      string
	ResetToken    string
	ResetTokenExp time.Time
	Active        bool
	Roles         UserRoles
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Constructor for creating a new user.
func NewUser(uid domain.UuidValueObject, name string, email domain.EmailValueObject, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:       uid,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Active:   true,
		Roles:    UserRoles{RoleUser},
	}
	user.Record(NewUserRegisteredEvent(uid.String(), user.Email.Value(), user.Roles))
	return user, nil
}

// Helper methods.
func (u *User) GetID() string {
	return u.ID.String()
}

// IsSuperAdmin checks if the user is a super admin.
func (u *User) IsSuperAdmin() bool {
	for _, role := range u.Roles {
		if role == RoleSuperAdmin {
			return true
		}
	}
	return false
}

// SetPassword sets the password for the user.
func (u *User) SetPassword(password string) {
	hashedPassword, _ := hashPassword(password)
	u.Password = hashedPassword
}

// hashPassword hashes the password using bcrypt.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// GenerateResetToken generates a random token for resetting the password.
func GenerateResetToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", token), nil
}
