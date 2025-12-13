package main

import (
	"sync"
)

// SimpleBroker - much simpler version with last 10 lines buffer
type SimpleBroker struct {
	clients   map[chan string]bool
	lastLines []string // Buffer to store last 10 lines
	mutex     sync.RWMutex
	maxLines  int // Maximum number of lines to keep (10)
}

func NewSimpleBroker() *SimpleBroker {
	return &SimpleBroker{
		clients:   make(map[chan string]bool),
		lastLines: make([]string, 0),
		maxLines:  10,
	}
}

// AddClient - add a new client and send last 10 lines
func (sb *SimpleBroker) AddClient(client chan string) {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()

	// Add client to map
	sb.clients[client] = true

	// Send last 10 lines to the new client
	go func() {
		sb.mutex.RLock()
		lines := make([]string, len(sb.lastLines))
		copy(lines, sb.lastLines)
		sb.mutex.RUnlock()

		// Send each historical line to the client
		for _, line := range lines {
			select {
			case client <- line:
				// Line sent successfully
			default:
				// Client channel is full, stop sending
				return
			}
		}
	}()
}

// RemoveClient - remove a client
func (sb *SimpleBroker) RemoveClient(client chan string) {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()
	delete(sb.clients, client)
	close(client)
}

// Broadcast - send message to all clients and store in buffer
func (sb *SimpleBroker) Broadcast(message string) {
	sb.mutex.Lock()

	// Add message to lastLines buffer
	sb.lastLines = append(sb.lastLines, message)

	// Keep only last 10 lines (remove oldest if more than 10)
	if len(sb.lastLines) > sb.maxLines {
		sb.lastLines = sb.lastLines[len(sb.lastLines)-sb.maxLines:]
	}

	// Get a copy of clients for broadcasting
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
			// Client channel is full or closed, remove it
			go sb.RemoveClient(client)
		}
	}
}
