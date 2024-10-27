package register

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"user-service/internal/user/domain"
	"user-service/kit"
	kitDomain "user-service/kit/domain"
	"user-service/kit/event"
)

type UserRegisterService struct {
	userRepository domain.UserRepository
	eventBus       event.Bus
}

func NewUserRegisterService(repo domain.UserRepository, eventBus event.Bus) *UserRegisterService {
	return &UserRegisterService{userRepository: repo, eventBus: eventBus}
}

// RegisterUser handles the registration logic
func (s *UserRegisterService) RegisterUser(id, name, email, password string) (*domain.User, *kit.DomainError) {
	uidValueObject, err := kitDomain.NewUuidValueObject(id)
	if err != nil {
		log.Println(err)
		return nil, domain.NewInvalidIDError()
	}
	emailValueObject, err := kitDomain.NewEmailValueObject(email)
	if err != nil {
		return nil, domain.NewInvalidEmailError()
	}

	existingUser, err := s.userRepository.FindByEmail(emailValueObject.Value())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			existingUser = nil
		} else {
			return nil, kit.NewDomainError(err.Error(), "user.register.database_error")
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
		return nil, kit.NewDomainError(err.Error(), "user.register.invalid_user")
	}
	if err := s.userRepository.Save(user); err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.register.database_error")
	}
	err = s.eventBus.Publish(user.PullEvents())
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.register.event_error")
	}
	return user, nil
}
