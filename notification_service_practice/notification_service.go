package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type NotificationRequest struct {
	requestID string
	Message   string
	UserId    string
	Channels  []string
}

type NotificationService struct {
	NotificationStatus       map[string]string // request_id-channel -> status
	mu                       sync.RWMutex
	RetryCount               int
	NotificationChannels     map[string]NotificationSender
	NotificationQueueChannel chan NotificationRequest
	StopNotificationChannel  chan struct{}
}

func (n *NotificationService) listenToAsyncRequest() {
	for {
		select {
		case notificationRequest := <-n.NotificationQueueChannel:
			n.sendNotificationUtil(&notificationRequest)
		case <-n.StopNotificationChannel:
			fmt.Println("stopping the consumer")
			return
		}
	}
}

func (n *NotificationService) sendNotificationUtil(req *NotificationRequest) {
	for _, channel := range req.Channels {
		errorCount := 0
		notificationStatusKey := fmt.Sprintf("%s-%s", req.requestID, channel)

		// Check status safely
		n.mu.RLock()
		status, ok := n.NotificationStatus[notificationStatusKey]
		n.mu.RUnlock()

		if ok && status == "success" {
			fmt.Println("Notification already sent")
			continue
		}

		// Mark in progress safely
		n.mu.Lock()
		n.NotificationStatus[notificationStatusKey] = "IN_PROGRESS"
		n.mu.Unlock()

		for errorCount < n.RetryCount {
			err := n.NotificationChannels[channel].SendNotification(req.requestID, req.UserId, req.Message)
			if err != nil {
				errorCount++
				time.Sleep(time.Second)
			} else {
				fmt.Println("logged success")

				n.mu.Lock()
				n.NotificationStatus[notificationStatusKey] = "success"
				n.mu.Unlock()
				break
			}
		}
		if errorCount == n.RetryCount {
			n.mu.Lock()
			n.NotificationStatus[notificationStatusKey] = "failed"
			n.mu.Unlock()
		}
	}
}

func NewNotificationService() *NotificationService {
	notificationService := &NotificationService{
		NotificationStatus: make(map[string]string),
		RetryCount:         3,
		NotificationChannels: map[string]NotificationSender{
			"email": &EmailSender{},
			"pn":    &PNSender{},
			"sms":   &SMSSender{},
		},
		NotificationQueueChannel: make(chan NotificationRequest, 10),
		StopNotificationChannel:  make(chan struct{}),
	}
	go notificationService.listenToAsyncRequest()
	return notificationService
}

func (n *NotificationService) SendNotification(req *NotificationRequest) {
	n.NotificationQueueChannel <- *req
}

func (n *NotificationService) GetNotificationStatus(request_id string, channel string) (string, error) {
	notificationKey := fmt.Sprintf("%s-%s", request_id, channel)
	n.mu.RLock()
	status, ok := n.NotificationStatus[notificationKey]
	n.mu.RUnlock()

	if !ok {
		return "", errors.New("requested notification not found")
	}
	return status, nil
}

func (n *NotificationService) LogAllNotification() {
	n.mu.RLock()
	defer n.mu.RUnlock()
	for key, val := range n.NotificationStatus {
		fmt.Println("notificationKey:", key, "notificationValue:", val)
	}
}

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

func main() {
	notificationService := NewNotificationService()

	notificationService.SendNotification(&NotificationRequest{
		requestID: "fasdfasdf",
		UserId:    "absece",
		Message:   "I love you",
		Channels:  []string{"email", "sms"},
	})

	// Give consumer time to process (simulating async)
	time.Sleep(2 * time.Second)

	notificationService.LogAllNotification()

	notificationStatus, err := notificationService.GetNotificationStatus("fasdfasdf", "email")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("notification status:", notificationStatus)

	// Stop consumer
	notificationService.StopNotificationChannel <- struct{}{}
	fmt.Println("consumer stopped")
}
