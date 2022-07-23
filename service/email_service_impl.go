package service

import (
	"gopkg.in/gomail.v2"
)

type EmailServiceImpl struct {
	Mailer *gomail.Message
	Dialer *gomail.Dialer
}

func NewEmailService(m *gomail.Message, d *gomail.Dialer) EmailService {
	return &EmailServiceImpl{
		Mailer: m,
		Dialer: d,
	}
}

func (es *EmailServiceImpl) SendEmail(message string) error {
	es.Mailer.SetHeader("From", CONFIG_SENDER_NAME)
	es.Mailer.SetHeader("To", CONFIG_AUTH_EMAIL)
	es.Mailer.SetHeader("Subject", "Product Created")
	es.Mailer.SetBody("text/html", message)

	err := es.Dialer.DialAndSend(es.Mailer)
	if err != nil {
		return err //500 Internal Server Error
	}
	es.Mailer.Reset()
	return nil
}
