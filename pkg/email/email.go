package email

import (
	"fmt"
	"go-shorten/config"
	"net/smtp"
)

type EmailService interface {
	SendVerificationEmail(to, subject, body string) error
}

type smtpEmailService struct {
	config *config.SMTPConfig
}

func NewSMTPSender(config *config.SMTPConfig) EmailService {
	return &smtpEmailService{config: config}
}

func (s *smtpEmailService) SendVerificationEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.config.Email, s.config.Password, s.config.Host)
	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		auth,
		s.config.FromAddress,
		[]string{to},
		msg,
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
