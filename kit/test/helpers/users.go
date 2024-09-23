package helpers

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"user-service/internal/user/domain"
	"user-service/kit/model"
)

func CreateUser() *domain.User {
	uid := uuid.New()
	baseUser := model.Base{
		ID: uid[:],
	}
	return &domain.User{
		Base:     baseUser,
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
		Roles:    []string{domain.RoleUser},
	}
}

func CreateManyUsers(n int) []*domain.User {
	users := make([]*domain.User, n)
	for i := 0; i < n; i++ {
		users[i] = CreateUser()
	}
	return users
}
