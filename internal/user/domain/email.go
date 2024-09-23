package domain

//go:generate mockery --case=snake --outpkg=domainmocks --output=domainmocks --name=EmailService

type EmailService interface {
	SendPasswordResetEmail(to, token string) error
	SendEmail(to, subject, body string) error
}
