package main

import (
	"sync"
	"time"
)

// RobustBroker - handles edge cases properly
type RobustBroker struct {
	clients     map[chan string]bool
	lastLines   []string
	mutex       sync.RWMutex
	maxLines    int
	clientCount int
	stats       *BrokerStats
}

type BrokerStats struct {
	TotalClients    int
	ActiveClients   int
	MessagesSent    int64
	MessagesDropped int64
	LastMessageTime time.Time
}

func NewRobustBroker() *RobustBroker {
	return &RobustBroker{
		clients:   make(map[chan string]bool),
		lastLines: make([]string, 0),
		maxLines:  10,
		stats:     &BrokerStats{},
	}
}

// AddClient - add client with proper error handling
func (rb *RobustBroker) AddClient(client chan string) bool {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	// Check if we have too many clients (prevent DoS)
	if len(rb.clients) >= 1000 {
		close(client)
		return false
	}

	rb.clients[client] = true
	rb.clientCount++
	rb.stats.TotalClients++
	rb.stats.ActiveClients++

	// Send historical lines with timeout
	go rb.sendHistoricalLines(client)

	return true
}

// sendHistoricalLines - send last 10 lines with timeout protection
func (rb *RobustBroker) sendHistoricalLines(client chan string) {
	rb.mutex.RLock()
	lines := make([]string, len(rb.lastLines))
	copy(lines, rb.lastLines)
	rb.mutex.RUnlock()

	// Send with timeout to prevent blocking
	timeout := time.After(5 * time.Second)

	for _, line := range lines {
		select {
		case client <- line:
			// Line sent successfully
		case <-timeout:
			// Timeout - client is too slow, remove it
			rb.RemoveClient(client)
			return
		}
	}
}

// RemoveClient - remove client safely
func (rb *RobustBroker) RemoveClient(client chan string) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	if _, exists := rb.clients[client]; exists {
		delete(rb.clients, client)
		rb.stats.ActiveClients--

		// Close channel safely
		select {
		case <-client:
			// Channel already closed
		default:
			close(client)
		}
	}
}

// Broadcast - send message with error handling and rate limiting
func (rb *RobustBroker) Broadcast(message string) {
	rb.mutex.Lock()

	// Add to buffer
	rb.lastLines = append(rb.lastLines, message)
	if len(rb.lastLines) > rb.maxLines {
		rb.lastLines = rb.lastLines[len(rb.lastLines)-rb.maxLines:]
	}

	// Update stats
	rb.stats.MessagesSent++
	rb.stats.LastMessageTime = time.Now()

	// Get clients copy
	clients := make(map[chan string]bool)
	for client := range rb.clients {
		clients[client] = true
	}

	rb.mutex.Unlock()

	// Broadcast with timeout and error handling
	rb.broadcastToClients(message, clients)
}

// broadcastToClients - broadcast with proper error handling
func (rb *RobustBroker) broadcastToClients(message string, clients map[chan string]bool) {
	timeout := time.After(2 * time.Second) // 2 second timeout per message

	for client := range clients {
		select {
		case client <- message:
			// Message sent successfully
		case <-timeout:
			// Client is too slow, remove it
			rb.RemoveClient(client)
			rb.mutex.Lock()
			rb.stats.MessagesDropped++
			rb.mutex.Unlock()
		default:
			// Client channel is full, remove it
			rb.RemoveClient(client)
			rb.mutex.Lock()
			rb.stats.MessagesDropped++
			rb.mutex.Unlock()
		}
	}
}

// GetStats - get broker statistics
func (rb *RobustBroker) GetStats() BrokerStats {
	rb.mutex.RLock()
	defer rb.mutex.RUnlock()

	stats := *rb.stats
	stats.ActiveClients = len(rb.clients)
	return stats
}

// Cleanup - cleanup disconnected clients
func (rb *RobustBroker) Cleanup() {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	// Remove clients with closed channels
	for client := range rb.clients {
		select {
		case <-client:
			// Channel is closed, remove it
			delete(rb.clients, client)
			rb.stats.ActiveClients--
		default:
			// Channel is still open
		}
	}
}
