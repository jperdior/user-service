package update_user

import (
	"user-service/internal/user/application/dto"
	"user-service/internal/user/domain"
	"user-service/kit"
	kitDomain "user-service/kit/domain"
)

type UpdateUserService struct {
	userRepository domain.UserRepository
}

func NewUpdateUserService(repo domain.UserRepository) *UpdateUserService {
	return &UpdateUserService{userRepository: repo}
}

func (s *UpdateUserService) UpdateUser(authenticatedUser *domain.AuthenticatedUser, ID, name, email, password string, roles []string) (*dto.UserDTO, error) {
	if !authenticatedUser.IsSuperAdmin() && authenticatedUser.ID != ID {
		return nil, domain.NewUnauthorizedError()
	}
	uid, err := kitDomain.NewUuidValueObject(ID)
	if err != nil {
		return nil, domain.NewInvalidIDError()
	}
	user, err := s.userRepository.FindByID(uid)
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.update_user.error")
	}
	if user == nil {
		return nil, domain.NewUserNotFoundError()
	}
	// Update only the non-empty fields
	if name != "" {
		user.Name = name
	}
	if email != "" {
		emailVO, err := kitDomain.NewEmailValueObject(email)
		if err != nil {
			return nil, domain.NewInvalidEmailError()
		}
		user.Email = emailVO
	}
	if password != "" {
		user.SetPassword(password)
	}
	if len(roles) > 0 {
		user.Roles = roles
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.update_user.error")
	}

	return dto.NewUserDTO(user), nil
}
