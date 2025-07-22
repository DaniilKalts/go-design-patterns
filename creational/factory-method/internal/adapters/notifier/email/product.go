package email

import (
	"github.com/DaniilKalts/go-design-patterns/factory-method/internal/adapters/notifier"
	"github.com/go-mail/mail"
)

var _ notifier.Notifier = (*EmailNotifier)(nil)

// EmailNotifier : Concrete Product
type EmailNotifier struct {
	dialer      *mail.Dialer
	fromAddress string
}

func (e *EmailNotifier) Send(to, subject, body string) error {
	msg := mail.NewMessage()

	msg.SetHeader("From", e.fromAddress)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	return e.dialer.DialAndSend(msg)
}
