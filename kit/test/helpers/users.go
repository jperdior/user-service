package helpers

import (
	"github.com/go-faker/faker/v4"
	"user-service/internal/user/domain"
	kitDomain "user-service/kit/domain"
)

func CreateUser() *domain.User {
	email, _ := kitDomain.NewEmailValueObject(faker.Email())
	return &domain.User{
		ID:       kitDomain.RandomUuidValueObject(),
		Name:     faker.Name(),
		Email:    email,
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
