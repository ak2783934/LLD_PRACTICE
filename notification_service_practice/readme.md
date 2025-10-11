Question: 

"Avinash, let’s design a Notification Service. Imagine we’re building a system that can send notifications to users through multiple channels — like Email, SMS, and Push Notifications.

We want this to be used by other internal systems (like an order service or payment service), which can trigger notifications to users based on events (like order placed, payment successful, etc.)."


🧩 Requirements:

Functional Requirements:
    The service should support multiple channels — Email, SMS, Push Notifications.
    It should allow sending notifications asynchronously.
    The system should handle failures gracefully (retry failed notifications).
    Support template-based notifications (e.g., message templates with placeholders like {username}, {order_id}).
    Allow adding new notification channels easily in the future (like WhatsApp, Slack, etc.).
    Provide an API for other services to trigger notifications.

Non-Functional Requirements:
    Scalability — the service should handle thousands of notifications per second.
    Reliability — ensure no notification is lost.
    Extensibility — new channels can be added with minimal changes.
    Observability — logs, metrics, and status tracking (like "SENT", "FAILED").



Assumptions: 
NotificationRequest{
    message string
    userID string 
    channel []string
}

NotificationService{
    
}
func(n *NotifiactionSErvfice)SendMessage()



NotificationSender interface{
    SendMessage(userID string, message string) error 
}

EmailNotificationSender 