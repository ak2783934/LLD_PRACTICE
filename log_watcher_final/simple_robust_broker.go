package main

import (
	"sync"
)

// SimpleRobustBroker - simple version with just essential edge case handling
type SimpleRobustBroker struct {
	clients   map[chan string]bool
	lastLines []string
	mutex     sync.RWMutex
	maxLines  int
}

func NewSimpleRobustBroker() *SimpleRobustBroker {
	return &SimpleRobustBroker{
		clients:   make(map[chan string]bool),
		lastLines: make([]string, 0),
		maxLines:  10,
	}
}

// AddClient - add client and send last 10 lines
func (sb *SimpleRobustBroker) AddClient(client chan string) {
	sb.mutex.Lock()
	sb.clients[client] = true
	sb.mutex.Unlock()

	// Send last 10 lines to new client
	go func() {
		sb.mutex.RLock()
		lines := make([]string, len(sb.lastLines))
		copy(lines, sb.lastLines)
		sb.mutex.RUnlock()

		for _, line := range lines {
			client <- line
		}
	}()
}

// RemoveClient - remove client safely
func (sb *SimpleRobustBroker) RemoveClient(client chan string) {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()

	if _, exists := sb.clients[client]; exists {
		delete(sb.clients, client)
		close(client)
	}
}

// Broadcast - send message to all clients
func (sb *SimpleRobustBroker) Broadcast(message string) {
	sb.mutex.Lock()

	// Add to buffer
	sb.lastLines = append(sb.lastLines, message)
	if len(sb.lastLines) > sb.maxLines {
		sb.lastLines = sb.lastLines[len(sb.lastLines)-sb.maxLines:]
	}

	// Get clients copy
	clients := make(map[chan string]bool)
	for client := range sb.clients {
		clients[client] = true
	}

	sb.mutex.Unlock()

	// Broadcast to all clients
	for client := range clients {
		select {
		case client <- message:
			// Message sent successfully
		default:
			// Client is slow, remove it
			sb.RemoveClient(client)
		}
	}
}
