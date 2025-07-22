package sms

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/DaniilKalts/go-design-patterns/factory-method/internal/adapters/notifier"
)

var _ notifier.Factory = (*SMSFactory)(nil)

// SMSNotifier : Concrete Product
type SMSNotifier struct {
	client    *twilio.RestClient
	fromPhone string
}

func (s *SMSNotifier) Send(to, subject, body string) error {
	msg := body
	if subject != "" {
		msg = fmt.Sprintf("%s\n\n%s", subject, msg)
	}

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(s.fromPhone)
	params.SetBody(msg)

	_, err := s.client.Api.CreateMessage(params)

	return err
}
