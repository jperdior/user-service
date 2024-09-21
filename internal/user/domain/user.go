package domain

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user-service/kit/model"
)

const (
	RoleUser       = "ROLE_USER"
	RoleSuperAdmin = "ROLE_SUPER_ADMIN"
)

type UserRoles []string

type User struct {
	model.Base
	model.BaseAggregate
	Name          string
	Email         string `gorm:"unique"`
	Password      string
	ResetToken    string
	ResetTokenExp time.Time
	Active        bool
	Roles         UserRoles `gorm:"serializer:json"`
}

func NewUser(uid model.UuidValueObject, name string, email model.EmailValueObject, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	base, _ := model.NewBase(uid)
	user := &User{
		Base:     *base,
		Name:     name,
		Email:    email.Value(),
		Password: hashedPassword,
		Active:   true,
		Roles:    UserRoles{RoleUser},
	}
	user.Record(NewUserRegisteredEvent(uid.String(), user.Email, user.Roles))
	return user, nil
}

func (u *User) GetID() string {
	uid, err := uuid.FromBytes(u.Base.ID)
	if err != nil {
		return ""
	}
	return uid.String()
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateResetToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", token), nil
}
