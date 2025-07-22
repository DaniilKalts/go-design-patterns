package sms

import (
	"github.com/twilio/twilio-go"

	"github.com/DaniilKalts/go-design-patterns/factory-method/internal/adapters/notifier"
)

var _ notifier.Factory = (*SMSFactory)(nil)

// SMSFactory : Concrete Creator
type SMSFactory struct {
	AccountSID string
	AuthToken  string
	FromPhone  string
}

func (f *SMSFactory) CreateNotifier() notifier.Notifier {
	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: f.AccountSID,
			Password: f.AuthToken,
		},
	)

	return &SMSNotifier{client: client, fromPhone: f.FromPhone}
}
