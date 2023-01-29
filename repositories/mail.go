package repositories

import (
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
)

func SendMail(recipient string, subject string, message string) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SENDER_NAME"))
	mailer.SetHeader("To", recipient)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", message)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("AUTH_EMAIL"),
		os.Getenv("AUTH_PASSWORD"),
		)
		
	err = dialer.DialAndSend(mailer)
	return err
}