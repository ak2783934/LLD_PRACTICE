package main

import "fmt"

type Notification interface {
	SendNotification(userID string, message string)
}

type EmailNotification struct{}

func (E *EmailNotification) SendNotification(userID string, message string) {
	fmt.Println("sending email notification")
}

type PushNotification struct{}

func (E *PushNotification) PushNotification(userID string, message string) {
	fmt.Println("sending push notification")
}
