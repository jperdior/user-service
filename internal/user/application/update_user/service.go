package update_user

import (
	"user-service/internal/user/application/dto"
	"user-service/internal/user/domain"
	"user-service/kit"
	"user-service/kit/model"
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
	uid, err := model.NewUuidValueObject(ID)
	if err != nil {
		return nil, domain.NewInvalidIDError()
	}
	user, err := s.userRepository.FindByID(uid.String())
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.update_user.error", 500)
	}
	if user == nil {
		return nil, domain.NewUserNotFoundError()
	}
	// Update only the non-empty fields
	if name != "" {
		user.Name = name
	}
	if email != "" {
		emailVO, err := model.NewEmailValueObject(email)
		if err != nil {
			return nil, domain.NewInvalidEmailError()
		}
		user.Email = emailVO.Value()
	}
	if password != "" {
		user.SetPassword(password)
	}
	if len(roles) > 0 {
		user.Roles = roles
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return nil, kit.NewDomainError(err.Error(), "user.update_user.error", 500)
	}

	return dto.NewUserDTO(user), nil
}
