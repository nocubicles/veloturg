package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmail(to []string, msg []byte) error {

	emailHost := os.Getenv("emailHost")
	smtpPort := os.Getenv("smtpPort")
	fromEmail := os.Getenv("fromEmail")
	smtpUser := os.Getenv("smtpUser")
	smtpPsw := os.Getenv("smtpPsw")

	auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPsw,
		emailHost,
	)

	err := smtp.SendMail(emailHost+":"+smtpPort, auth, fromEmail, to, msg)

	if err != nil {
		return err
	}
	return nil
}

func SendSignInEmail(to string, signInToken string) error {
	server := os.Getenv("SERVER")
	loginLink := fmt.Sprintf("<a href='%s/signin?signintoken=%s'>lingile</a>", server, signInToken)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: Teie sisse logimise link veloturg.ee\n"+
			mime+
			"<html><body>"+
			"<p>Palun vajutage %s, et logida sisse.</p>"+
			"<p>veloturg.ee</p>"+
			"</body></html>",
		to, loginLink,
	)

	err := sendEmail([]string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
