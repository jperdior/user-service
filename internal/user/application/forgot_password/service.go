package forgot_password

import (
	"user-service/internal/user/domain"
	"user-service/kit"
)

type ForgotPasswordService struct {
	userRepository domain.UserRepository
	mailer         domain.EmailService
}

func NewForgotPasswordService(repo domain.UserRepository, mailer domain.EmailService) *ForgotPasswordService {
	return &ForgotPasswordService{userRepository: repo, mailer: mailer}
}

func (s *ForgotPasswordService) SendResetPasswordEmail(email string) *kit.DomainError {
	_, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return domain.NewUserNotFoundError()
	}

	// send email with password reset link
	err = s.mailer.SendPasswordResetEmail(email)
	if err != nil {
		return kit.NewDomainError(err.Error(), "user.forgot_password.email_error", 500)
	}
	return nil
}
