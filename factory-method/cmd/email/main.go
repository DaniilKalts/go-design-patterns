package main

import (
	"log"
	"os"

	"github.com/DaniilKalts/go-design-patterns/factory-method/internal/adapters/notifier/email"
)

func main() {
	factory := &email.EmailFactory{
		SMTPServer:  os.Getenv("SMTP_SERVER"),
		SMTPPort:    587,
		Username:    os.Getenv("SMTP_USER"),
		Password:    os.Getenv("SMTP_PASS"),
		FromAddress: os.Getenv("EMAIL_FROM"),
	}

	notifier := factory.CreateNotifier()
	if err := notifier.Send(
		os.Getenv("TO_EMAIL"),
		"Gopher learning Design Patterns!",
		"Join us! Its thrilling!",
	); err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully")
}
