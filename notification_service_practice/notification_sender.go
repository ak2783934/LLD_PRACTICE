package main

import "fmt"

type NotificationSender interface {
	SendNotification(requestID string, userId string, message string) error
}

type EmailSender struct{}

func (e *EmailSender) SendNotification(requestID string, userId string, message string) error {
	fmt.Println("Email sent to user id:", userId, "message:", message)
	return nil
}

type SMSSender struct{}

func (e *SMSSender) SendNotification(requestID string, userId string, message string) error {
	fmt.Println("SMS sent to user id:", userId, "message:", message)
	return nil
}

type PNSender struct{}

func (p *PNSender) SendNotification(requestID string, userId string, message string) error {
	fmt.Println("PN sent to user id:", userId, "message:", message)
	return nil
}
