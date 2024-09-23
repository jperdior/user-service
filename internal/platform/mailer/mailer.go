package mailer

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type MailerConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Mailer struct {
	configuration MailerConfig
}

var MAILER *gomail.Dialer

func NewMailer(configuration MailerConfig) {
	dialer := gomail.NewDialer(configuration.Host, configuration.Port, configuration.User, configuration.Password)
	if configuration.Port == 465 {
		dialer.SSL = true
	}
	// TODO: improve this
	if configuration.Port == 587 {
		dialer.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // This is not recommended for production
		}
	}
	MAILER = dialer
}
