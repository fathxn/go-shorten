package email

import "net/smtp"

type Sender interface {
	SendVerificationEmail(to, verificationLink string) error
}

type smtpSender struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

func NewSMTPSender(from, password, smtpHost, smtpPort string) Sender {
	return &smtpSender{
		from:     from,
		password: password,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
	}
}

func (s *smtpSender) SendVerificationEmail(to, verificationLink string) error {
	subject := "Ringkas.Link - Email Verification"
	body := "Please click the following link to verify your email: " + verificationLink

	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", s.from, s.password, s.smtpHost)
	return smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.from, []string{to}, message)
}
