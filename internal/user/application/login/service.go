package login

import (
	"golang.org/x/crypto/bcrypt"
	"user-service/internal/user/domain"
	"user-service/kit"
)

type UserLoginService struct {
	userRepository domain.UserRepository
	jwtService     domain.TokenService
}

func NewUserLoginService(repo domain.UserRepository, tokenService domain.TokenService) *UserLoginService {
	return &UserLoginService{userRepository: repo, jwtService: tokenService}
}

func (s *UserLoginService) Login(email, password string) (string, *kit.DomainError) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", domain.NewInvalidCredentialsError()
	}
	if !checkPasswordHash(password, user.Password) {
		return "", domain.NewInvalidCredentialsError()
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return "", domain.NewInvalidCredentialsError()
	}

	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
