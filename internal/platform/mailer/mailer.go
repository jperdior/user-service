package mailer

import "gopkg.in/gomail.v2"

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
	MAILER = dialer
}
