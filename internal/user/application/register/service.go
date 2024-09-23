package register

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"user-service/internal/user/domain"
	"user-service/kit"
	"user-service/kit/model"
)

type UserRegisterService struct {
	userRepository domain.UserRepository
}

func NewUserRegisterService(repo domain.UserRepository) *UserRegisterService {
	return &UserRegisterService{userRepository: repo}
}

// RegisterUser handles the registration logic
func (s *UserRegisterService) RegisterUser(id, name, email, password string) (*domain.User, *kit.DomainError) {
	uidValueObject, err := model.NewUuidValueObject(id)
	if err != nil {
		log.Println(err)
		return nil, domain.NewInvalidIDError()
	}
	emailValueObject, err := model.NewEmailValueObject(email)
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
