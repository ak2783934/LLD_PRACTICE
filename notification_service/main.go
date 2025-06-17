package main

import "fmt"

type Notification struct {
	// Add fields here if needed
}

func (n *Notification) Send(to, message string) error {
	// Implement the send logic here
	return nil
}

type EmailSender struct{}

func (e *EmailSender) Send(to, message string) error {
	fmt.Printf("Email sent to %s: %s\n", to, message)
	return nil
}

type SMSSender struct{}

func (s *SMSSender) Send(to, message string) error {
	fmt.Printf("SMS sent to %s: %s\n", to, message)
	return nil
}

type PushSender struct{}

func (p *PushSender) Send(to, message string) error {
	fmt.Printf("Push sent to %s: %s\n", to, message)
	return nil
}

type NotificationService struct {
	senders map[string]Notification
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		senders: map[string]Notification{
			"email": EmailSender{},
			"push":  PushSender{},
			"sms":   SMSSender{},
		},
	}
}

func (ns *NotificationService) Send(channel, to, message string) {
	sender, ok := ns.senders[channel]
	if !ok {
		return fmt.Errorf("unsupported channel: %s", channel)
	}
	return sender.Send(to, message)
}

func main() {
	ns := NewNotificationService()

	ns.Send("email", "alice@example.com", "Welcome to Safe Security!")
	ns.Send("sms", "+919000000000", "OTP is 123456")
	ns.Send("push", "user123", "New vulnerability detected")
}
