package application

import "user-service/internal/user/domain"

type UserService struct {
	repo domain.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser use case
func (s *UserService) CreateUser(name, email, password string) (*domain.User, error) {
	user, err := domain.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}
