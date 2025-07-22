package main

import (
	"log"
	"os"

	"github.com/DaniilKalts/go-design-patterns/factory-method/internal/adapters/notifier/sms"
)

func main() {
	factory := &sms.SMSFactory{
		AccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		FromPhone:  os.Getenv("TWILIO_FROM_PHONE"),
	}

	notifier := factory.CreateNotifier()
	if err := notifier.Send(
		os.Getenv("TO_PHONE"),
		"Gopher learning Design Patterns!",
		"Join us! Its thrilling!",
	); err != nil {
		log.Fatal(err)
	}

	log.Println("SMS sent successfully")
}
