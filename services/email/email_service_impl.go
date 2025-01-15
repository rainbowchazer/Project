package email

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"text/template"
)

type EmailService struct {
	SMTPHost     string
	SMTPPort     string
	SenderEmail  string
	SenderPass   string
	TemplateFS   embed.FS
	ContextMutex sync.Mutex
}

func NewEmailService(smtpHost, smtpPort, senderEmail, senderPass string, templateFS embed.FS) *EmailService {
	return &EmailService{
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SenderEmail:  senderEmail,
		SenderPass:   senderPass,
		TemplateFS:   templateFS,
		ContextMutex: sync.Mutex{},
	}
}

func (s *EmailService) SendEmailAsync(ctx context.Context, emailContext *EmailContext, receiver string) {
	go func() {
		if err := s.SendEmail(emailContext, receiver); err != nil {
			log.Printf("Error sending email: %v\n", err)
		}
	}()
}

func (s *EmailService) SendEmail(emailContext *EmailContext, receiver string) error {
	s.ContextMutex.Lock()
	defer s.ContextMutex.Unlock()

	templateData, err := s.TemplateFS.ReadFile(emailContext.TemplateLocation)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	tmpl, err := template.New("email").Parse(string(templateData))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	emailContext.Context["currentReceiver"] = receiver
	if err := tmpl.Execute(&body, emailContext.Context); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	auth := smtp.PlainAuth("", s.SenderEmail, s.SenderPass, s.SMTPHost)
	to := []string{receiver}
	msg := []byte(fmt.Sprintf("Subject: %s\nContent-Type: text/html; charset=UTF-8\n\n%s", emailContext.Subject, body.String()))

	if err := smtp.SendMail(fmt.Sprintf("%s:%s", s.SMTPHost, s.SMTPPort), auth, emailContext.Sender, to, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent to %s\n", receiver)
	return nil
}

func (s *EmailService) SendEmailToAll(emailContext *EmailContext) {
	for _, receiver := range emailContext.Receivers {
		s.SendEmailAsync(context.Background(), emailContext, receiver)
	}
}
