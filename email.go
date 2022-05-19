package main

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

func Mail(fromMail, password, toMail, subject, body, host string, port int) (string, error) {

	m := gomail.NewMessage()
	// Set E-Mail sender
	m.SetHeader("From", fromMail)

	// Set E-Mail receivers
	m.SetHeader("To", toMail)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer(host, port, fromMail, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return "Fail", err
	}

	return "Success", nil
}
