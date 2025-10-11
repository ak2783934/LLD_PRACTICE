package main

import (
	"errors"
	"fmt"
	"time"
)

type NotificationRequest struct {
	requestID string
	Message   string
	UserId    string
	Channels  []string
}

type NotificationService struct {
	NotificationStatus   map[string]string // request_id+-+channel -> status?
	RetryCount           int
	NotificationChannels map[string]NotificationSender
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		NotificationStatus: make(map[string]string),
		RetryCount:         3,
		NotificationChannels: map[string]NotificationSender{
			"email": &EmailSender{},
			"pn":    &PNSender{},
			"sms":   &SMSSender{},
		},
	}
}

// assuming things are sync
func (n *NotificationService) SendNotification(req *NotificationRequest) {
	channels := req.Channels

	for _, channel := range channels {
		errorCount := 0
		notificationStatusKey := req.requestID + "-" + channel
		status, ok := n.NotificationStatus[notificationStatusKey]
		if ok && status == "success" {
			fmt.Println("Notification already sent")
			continue
		}
		n.NotificationStatus[notificationStatusKey] = "IN_PROGRESS"

		for errorCount < n.RetryCount {
			err := n.NotificationChannels[channel].SendNotification(req.requestID, req.UserId, req.Message)
			if err != nil {
				errorCount++
				time.Sleep(time.Second)
			} else {
				fmt.Println("logged success")
				n.NotificationStatus[notificationStatusKey] = "success"
				break
			}
		}
		if errorCount == n.RetryCount {
			n.NotificationStatus[notificationStatusKey] = "failed"
		}
	}
}

func (n *NotificationService) GetNotificationStatus(request_id string, channel string) (string, error) {
	notificationKey := request_id + "-" + channel
	fmt.Println(notificationKey)
	status, ok := n.NotificationStatus[notificationKey]
	if !ok {
		return "", errors.New("requested notification not found")
	}
	return status, nil
}

func (n *NotificationService) LogAllNotification() {
	for key, val := range n.NotificationStatus {
		fmt.Println("notificationKey : ", key, " notificationValue", val)
	}
}

type NotificationSender interface {
	SendNotification(requestID string, userId string, message string) error
}

type EmailSender struct{}

func (e *EmailSender) SendNotification(requestID string, userId string, message string) error {

	fmt.Println("Email sent to user id: ", userId, " message ", message)
	return nil
}

type SMSSender struct{}

func (e *SMSSender) SendNotification(requestID string, userId string, message string) error {
	fmt.Println("SMS sent to user id: ", userId, " message ", message)
	return nil
}

type PNSender struct{}

func (p *PNSender) SendNotification(requestID string, userId string, message string) error {
	fmt.Println("PN sent to user id: ", userId, " message ", message)
	return nil
}

func main() {
	notificationService := NewNotificationService()

	notificationService.SendNotification(&NotificationRequest{
		requestID: "fasdfasdf",
		UserId:    "absece",
		Message:   "I love you",
		Channels:  []string{"email", "sms"},
	})

	notificationService.LogAllNotification()

	notificationStatus, err := notificationService.GetNotificationStatus("fasdfasdf", "email")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("notification status ", notificationStatus)

}
