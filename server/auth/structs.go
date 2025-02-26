package auth

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
)

type EMAILSENDER struct {
	conf config.EmailSenderConfig
	Host string
	Port string
}

func NewEmailSender(conf config.EmailSenderConfig, host string, port string) *EMAILSENDER {
	return &EMAILSENDER{
		conf: conf,
		Host: host,
		Port: port,
	}
}

// Using for Testing purpose
func TestEmailSender() (*EMAILSENDER, error) {

	conf, err := config.CreateNewEmailSender(
		"On-Sight Test",
		"TEST_EMAIL_SENDER",
		"TEST_EMAIL_PASSWORD",
	)

	if err != nil {
		return nil, err
	}

	host := "smtp.gmail.com"
	port := ":587"

	return NewEmailSender(*conf, host, port), nil
}

type EMAIL struct {
	Title string
	To    string
	MIME  string
	Body  string
}
