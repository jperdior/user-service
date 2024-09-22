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

func (u *UserRepositoryImpl) Find(criteria domain.Criteria) ([]*domain.User, int64, error) {
	var users []*domain.User
	db := u.db

	for _, filter := range criteria.Filters {
		if validFilter, ok := domain.ValidFilters[filter.Name]; ok {
			switch validFilter.Operation {
			case "LIKE":
				db = db.Where(validFilter.Name+" LIKE ?", "%"+filter.Value.(string)+"%")
			case "=":
				db = db.Where(validFilter.Name+" = ?", filter.Value)
			}
		}
	}

	var totalRows int64
	db.Model(&domain.User{}).Count(&totalRows)

	if criteria.Page > 0 {
		offset := (criteria.Page - 1) * criteria.PageSize
		db = db.Offset(offset)
	}
	if criteria.PageSize > 0 {
		db = db.Limit(criteria.PageSize)
	}
	if criteria.Sort != "" {
		db = db.Order(criteria.Sort)
	}
	err := db.Find(&users).Error
	return users, totalRows, err
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepositoryImpl) Save(user *domain.User) error {
	return u.db.Save(user).Error
}
