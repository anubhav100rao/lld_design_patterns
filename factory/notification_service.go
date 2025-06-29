package factory

import "fmt"

// Notifier interface
type Notifier interface {
	Send(to, message string) error
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(to, msg string) error {
	fmt.Printf("Sending EMAIL to %s: %s\n", to, msg)
	return nil
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(to, msg string) error {
	fmt.Printf("Sending SMS to %s: %s\n", to, msg)
	return nil
}

// Factory
func NewNotifier(channel string) (Notifier, error) {
	switch channel {
	case "email":
		return &EmailNotifier{}, nil
	case "sms":
		return &SMSNotifier{}, nil
	default:
		return nil, fmt.Errorf("unsupported channel: %s", channel)
	}
}
