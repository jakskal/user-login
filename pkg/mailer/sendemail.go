package mailer

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

// SendEmail sending email to receipent
func SendEmail(receipent string, content string, subject string) error {
	configSMTPHost := os.Getenv("CONFIG_SMTP_HOST")
	configSMTPPort := 587
	configEmail := os.Getenv("CONFIG_EMAIL")
	configPassword := os.Getenv("CONFIG_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", configEmail)
	mailer.SetHeader("To", receipent)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", content)

	dialer := gomail.NewDialer(
		configSMTPHost,
		configSMTPPort,
		configEmail,
		configPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	log.Printf("Send mail to %v", receipent)
	log.Println("Mail sent!")
	return nil
}
