package email

import (
	"github.com/DaniilKalts/go-design-patterns/factory-method/internal/adapters/notifier"
	"github.com/go-mail/mail"
)

// EmailFactory : Concrete Creator
type EmailFactory struct {
	SMTPServer  string
	SMTPPort    int
	Username    string
	Password    string
	FromAddress string
}

func (f *EmailFactory) CreateNotifier() notifier.Notifier {
	dialer := mail.NewDialer(
		f.SMTPServer,
		f.SMTPPort,
		f.Username,
		f.Password,
	)

	return &EmailNotifier{
		dialer:      dialer,
		fromAddress: f.FromAddress,
	}
}
