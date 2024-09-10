package domain

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user-service/kit/model"
)

type User struct {
	model.Base
	model.BaseAggregate
	Name          string
	Email         string `gorm:"unique"`
	Password      string
	ResetToken    string
	ResetTokenExp time.Time
	Active        bool
	Tokens        []UserToken
}

type UserToken struct {
	model.Base
	UserID    []byte `gorm:"type:binary(16);index"` // Foreign key to the User
	Token     string `gorm:"unique"`                // JWT token
	Device    string // Optional: to track which device the token belongs to
	ExpiresAt time.Time
}

// todo: MOVE TO A SERVICE

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateResetToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", token), nil
}
