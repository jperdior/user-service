package find_user

import (
	"user-service/internal/user/application/dto"
	"user-service/internal/user/domain"
	"user-service/kit"
	kitDomain "user-service/kit/domain"
)

type UserFinderService struct {
	userRepository domain.UserRepository
}

func NewUserFinderService(repo domain.UserRepository) *UserFinderService {
	return &UserFinderService{userRepository: repo}
}

func (s *UserFinderService) FindUser(authenticatedUser *domain.AuthenticatedUser, ID string) (*dto.UserDTO, *kit.DomainError) {
	if !authenticatedUser.IsSuperAdmin() && authenticatedUser.ID != ID {
		return nil, domain.NewUnauthorizedError()
	}
	uid, err := kitDomain.NewUuidValueObject(ID)
	if err != nil {
		return nil, domain.NewInvalidIDError()
	}
	user, err := s.userRepository.FindByID(uid.String())
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.find_user.error", 500)
	}
	if user == nil {
		return nil, domain.NewUserNotFoundError()
	}

	return dto.NewUserDTO(user), nil
}
