package domain

type EmailService interface {
	SendPasswordResetEmail(to string) error
	SendEmail(to, subject, body string) error
}
