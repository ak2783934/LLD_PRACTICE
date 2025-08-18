package main

import "fmt"

type Notification interface {
	sendMessage(string) string
}

type EmailNotification struct{}

func (e *EmailNotification) sendMessage(message string) string {
	return "email message " + message
}

type PushNotification struct{}

func (p *PushNotification) sendMessage(message string) string {
	return "push notification " + message
}

func NotificationFactory(method string) Notification {
	switch method {
	case "email":
		return &EmailNotification{}
	case "push":
		return &PushNotification{}
	}
	return nil
}

func main() {
	notifier := NotificationFactory("email")
	fmt.Println(notifier.sendMessage("Hello!"))

	notifier1 := NotificationFactory("push")
	fmt.Println(notifier1.sendMessage("Hi!"))
}
