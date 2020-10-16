package services

import(
	"crypto/tls"
	mail "gopkg.in/mail.v2"
)

func SendEmail(mailBody, toEmail, subject string) bool{
	emailSender := env["email_sender"]
	emailPassword := env["email_password"]
	

	msg := mail.NewMessage()
	msg.SetHeader("From", emailSender)
	msg.SetHeader("To", toEmail)
    msg.SetHeader("Subject", subject)
    msg.SetBody("text/plain", mailBody)
    d := mail.NewDialer("smtp.gmail.com", 587,emailSender, emailPassword)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(msg); err != nil {
      logger.Error.Println(err)
      return false
    }
  return true
}