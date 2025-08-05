package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/PrinceNarteh/social/internal/config"
)

type SendGridMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}

var SendGridClient *SendGridMailer

func NewSendGrid() {
	client := sendgrid.NewSendClient(config.Envs.MailConfig.ApiKey)
	SendGridClient = &SendGridMailer{
		apiKey:    config.Envs.MailConfig.ApiKey,
		fromEmail: config.Envs.MailConfig.FromEmail,
		client:    client,
	}
}

func (m *SendGridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	tmpl, err := template.ParseFS(FS, fmt.Sprintf("templates/%s", UserWelcomeTemplate))
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())
	message.SetMailSettings(
		&mail.MailSettings{
			SandboxMode: &mail.Setting{
				Enable: &isSandbox,
			},
		},
	)

	for i := range maxRetries {
		res, err := m.client.Send(message)
		if err != nil {
			log.Printf("failed to send email to %v, attempt %d of %d", email, i, maxRetries)
			log.Printf("Error: %v", err)
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		log.Printf("Email send with status code %v", res.StatusCode)
		return nil
	}

	return fmt.Errorf("failed to send email with %d attempts", maxRetries)
}
