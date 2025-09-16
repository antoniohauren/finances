package mailer

import (
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
)

func SendEmail(to []string, subject string, body string) error {
	smtpFrom := "finances@detrouse.com"
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	message := []byte(
		"From: " + smtpFrom + "\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			body + "\r\n\r\n",
	)

	err := smtp.SendMail(addr, nil, smtpFrom, to, message)

	if err != nil {
		slog.Error("send email", "error", err)
		return err
	}

	slog.Info("email sent successfully!")

	return nil
}
