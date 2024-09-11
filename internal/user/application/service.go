package application

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"user-service/internal/user/domain"
	"user-service/kit"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

// RegisterUser handles the registration logic
func (s *UserService) RegisterUser(id, name, email, password string) (*domain.User, *kit.DomainError) {
	uidValueObject, err := kit.NewUuidValueObject(id)
	if err != nil {
		log.Println(err)
		return nil, domain.NewInvalidIDError()
	}
	emailValueObject, err := kit.NewEmailValueObject(email)
	if err != nil {
		return nil, domain.NewInvalidEmailError()
	}

	existingUser, err := s.userRepository.FindByEmail(emailValueObject.Value())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			existingUser = nil
		} else {
			return nil, kit.NewDomainError(err.Error(), "user.register.database_error", http.StatusInternalServerError)
		}
	}
	if existingUser != nil {
		return nil, domain.NewExistingUserError()
	}

	user, err := domain.NewUser(
		uidValueObject,
		name,
		emailValueObject,
		password,
	)
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.register.invalid_user", http.StatusBadRequest)
	}
	if err := s.userRepository.Save(user); err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.register.database_error", http.StatusInternalServerError)
	}
	return user, nil
}
