package auth

import (
	"fmt"
	"net/smtp"
)

func (es *EMAILSENDER) PLAINAUTH() smtp.Auth {
	return smtp.PlainAuth(
		es.conf.Identity,
		es.conf.EMAIL,
		es.conf.PASSWORD,
		es.Host,
	)
}

func (es *EMAILSENDER) SendMail(to []string, message []byte) error {
	return smtp.SendMail(
		es.Host+es.Port,
		es.PLAINAUTH(),
		es.conf.EMAIL,
		to,
		message,
	)
}

func (es *EMAILSENDER) SendOTP(email string, otp int) error {

	to := []string{email}

	to_message := fmt.Sprintf("To: %s\r\n", to)

	subject_message := "Subject: OTP VERIFY\n"
	headers := "MIME version: 1.0;\nContent-Type: text/html;\ncharset:\"UTF-8\"\n\n"

	body_message := fmt.Sprintf(`
  <html>
  <body>
    <p> Please Enter the given otp code on the main page to complete registration </p><br>
   <h2><b> %d </b></h2> <br>
    <p> You can ignore this message if not requested by you</p><br>

  </body>
  </html>
  `, otp)

	rawMessage := []byte(to_message + subject_message + headers + body_message)

	return es.SendMail(to, rawMessage)
}
