package notifier

// Notifier : Product
type Notifier interface {
	Send(to, subject, body string) error
}

// Factory : Creator
type Factory interface {
	CreateNotifier() Notifier
}
