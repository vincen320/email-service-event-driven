package service

import "email-service-event-driven/app"

var env = app.NewViper()

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Email-Service-Event-Driven <youraddress@gmail.com>"

var CONFIG_AUTH_EMAIL = env.GetString("EMAIL_ADDRESS")
var CONFIG_AUTH_PASSWORD = env.GetString("EMAIL_PASSWORD")

type EmailService interface {
	SendEmail(message string) error
}
