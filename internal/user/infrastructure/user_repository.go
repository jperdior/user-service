package infrastructure

import (
	"gorm.io/gorm"
	"user-service/internal/user/domain"
	"user-service/kit"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) FindByID(id string) (*domain.User, error) {
	uuidBinary, err := kit.UuidStringToBinary(id)
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = u.db.Where("id = ?", uuidBinary).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepositoryImpl) Save(user *domain.User) error {
	return u.db.Save(user).Error
}
