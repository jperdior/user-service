package infrastructure

import "gopkg.in/gomail.v2"

type MailerImpl struct {
	dialer *gomail.Dialer
}

func NewMailer(dialer *gomail.Dialer) *MailerImpl {
	return &MailerImpl{dialer: dialer}
}

func (m *MailerImpl) SendEmail(to, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.dialer.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	return m.dialer.DialAndSend(message)
}
