package main

import (
	"testing"
)

func TestMail(t *testing.T) {
	fromMail := "abcdef@gmail.com"
	password := "ghijklamn"

	toMail := "abcdef@gmail.com"
	subject := "Gomail Mail Testing"
	body := "This is Gomail test body"
	host := "smtp.gmail.com"
	port := 587

	s, err := Mail(fromMail, password, toMail, subject, body, host, port)

	if s != "Success" && err != nil {
		t.Errorf("Mail not sent!!")
	}
}
