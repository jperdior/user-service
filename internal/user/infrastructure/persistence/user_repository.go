package persistence

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
	"user-service/internal/user/domain"
	kitDomain "user-service/kit/domain"
	"user-service/kit/infrastructure/persistence"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user *domain.User) error {
	userModel := userToUserModel(user)
	return r.db.Save(userModel).Error
}

func (r *GormUserRepository) Find(criteria domain.Criteria) ([]*domain.User, int64, error) {
	var users []*UserModel
	var domainUsers []*domain.User
	db := r.db

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
	db.Model(&UserModel{}).Count(&totalRows)

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
	for _, userModel := range users {
		domainUsers = append(domainUsers, userModelToDomainUser(userModel))
	}
	return domainUsers, totalRows, err
}

func (r *GormUserRepository) FindByID(id kitDomain.UuidValueObject) (*domain.User, error) {
	var userModel UserModel
	err := r.db.First(&userModel, "id = ?", id.Bytes()).Error
	if err != nil {
		return nil, err
	}
	return userModelToDomainUser(&userModel), nil
}

func (r *GormUserRepository) FindByEmail(email string) (*domain.User, error) {
	var userModel UserModel
	err := r.db.First(&userModel, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return userModelToDomainUser(&userModel), nil
}

func (r *GormUserRepository) Delete(user *domain.User) error {
	userModel := userToUserModel(user)
	return r.db.Delete(userModel).Error
}

type UserModel struct {
	persistence.Base
	Name          string
	Email         string `gorm:"unique"`
	Password      string
	ResetToken    string
	ResetTokenExp time.Time
	Active        bool
	Roles         []byte `gorm:"serializer:json"`
}

func (UserModel) TableName() string {
	return "user"
}

// userToUserModel converts domain.User to UserModel (DTO).
func userToUserModel(user *domain.User) *UserModel {
	base := persistence.Base{
		ID:        user.ID.Bytes(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	rolesJson, _ := json.Marshal(user.Roles)
	return &UserModel{
		Base:          base,
		Name:          user.Name,
		Email:         user.Email.Value(),
		Password:      user.Password,
		ResetToken:    user.ResetToken,
		ResetTokenExp: user.ResetTokenExp,
		Active:        user.Active,
		Roles:         rolesJson,
	}
}

// userModelToDomainUser converts UserModel (DTO) to domain.User.
func userModelToDomainUser(userModel *UserModel) *domain.User {
	uid, _ := kitDomain.NewUuidValueObjectFromBytes(userModel.ID)
	email, _ := kitDomain.NewEmailValueObject(userModel.Email)
	var roles []string
	_ = json.Unmarshal(userModel.Roles, &roles)
	user := &domain.User{
		BaseAggregate: kitDomain.BaseAggregate{},
		ID:            uid,
		Name:          userModel.Name,
		Email:         email,
		Password:      userModel.Password,
		ResetToken:    userModel.ResetToken,
		ResetTokenExp: userModel.ResetTokenExp,
		Active:        userModel.Active,
		Roles:         roles,
		CreatedAt:     userModel.CreatedAt,
		UpdatedAt:     userModel.UpdatedAt,
	}
	return user
}
