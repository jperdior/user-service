package domain

type EmailProvider interface {
	SendEmail(to, subject, body string) error
}
