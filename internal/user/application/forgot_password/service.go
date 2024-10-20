package forgot_password

import (
	"user-service/internal/user/domain"
	"user-service/kit"
)

type ForgotPasswordService struct {
	userRepository domain.UserRepository
	mailer         domain.EmailService
	tokenService   domain.TokenService
}

func NewForgotPasswordService(repo domain.UserRepository, mailer domain.EmailService, tokenService domain.TokenService) *ForgotPasswordService {
	return &ForgotPasswordService{userRepository: repo, mailer: mailer, tokenService: tokenService}
}

func (s *ForgotPasswordService) SendResetPasswordEmail(email string) *kit.DomainError {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return domain.NewUserNotFoundError()
	}

	token, err := s.tokenService.GenerateResetPasswordToken(user)

	// send email with password reset link
	err = s.mailer.SendPasswordResetEmail(user.Email.Value(), token)
	if err != nil {
		return kit.NewDomainError(err.Error(), "user.forgot_password.email_error", 500)
	}
	return nil
}
