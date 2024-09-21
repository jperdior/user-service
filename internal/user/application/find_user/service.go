package find_user

import (
	"user-service/internal/user/domain"
	"user-service/kit"
)

type UserFinderService struct {
	userRepository domain.UserRepository
}

func NewUserFinderService(repo domain.UserRepository) UserFinderService {
	return UserFinderService{userRepository: repo}
}

func (s UserFinderService) FindUser(ID string) (*domain.User, *kit.DomainError) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.find_user.error", 500)
	}
	if user == nil {
		return nil, domain.NewUserNotFoundError()
	}

	return user, nil
}
