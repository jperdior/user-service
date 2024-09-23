package infrastructure

import "gopkg.in/gomail.v2"

type EmailServiceImpl struct {
	dialer *gomail.Dialer
}

func NewEmailServiceImpl(dialer *gomail.Dialer) *EmailServiceImpl {
	return &EmailServiceImpl{dialer: dialer}
}

func (m *EmailServiceImpl) SendPasswordResetEmail(to, token string) error {
	body := "Click <a href='http://localhost:8080/reset-password?token=" + token + "'>here</a> to reset your password"
	return m.SendEmail(to, "Password Reset", body)
}

func (m *EmailServiceImpl) SendEmail(to, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.dialer.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	return m.dialer.DialAndSend(message)
}
