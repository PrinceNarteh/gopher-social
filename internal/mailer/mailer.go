package mailer

import "embed"

const (
	FromName            = "Gopher Social"
	maxRetries          = 3
	UserWelcomeTemplate = "user_inviation.templ"
)

//go:embed "templates"
var FS embed.FS

type MailerClient interface {
	Send(templateFile, username, email string, data any, isSandbox bool)
}
