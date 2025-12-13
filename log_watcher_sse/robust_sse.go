package main

import (
	"fmt"
	"net/http"
	"time"
)

// RobustSSEHandler - handles network failures and client disconnections
func RobustSSEHandler(broker *RobustBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

		// Check if client supports streaming
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		// Create client channel with buffer to prevent blocking
		clientChan := make(chan string, 100) // Buffer for 100 messages

		// Register client with timeout
		if !broker.AddClient(clientChan) {
			http.Error(w, "Too many clients", http.StatusServiceUnavailable)
			return
		}

		// Ensure cleanup on exit
		defer func() {
			broker.RemoveClient(clientChan)
		}()

		// Send initial connection message
		fmt.Fprintf(w, "data: Connected to log stream\n\n")
		flusher.Flush()

		// Heartbeat ticker to detect dead connections
		heartbeat := time.NewTicker(30 * time.Second)
		defer heartbeat.Stop()

		// Message handling loop
		for {
			select {
			case message, ok := <-clientChan:
				if !ok {
					// Channel closed, client disconnected
					return
				}

				// Send message to browser
				fmt.Fprintf(w, "data: %s\n\n", message)
				flusher.Flush()

			case <-heartbeat.C:
				// Send heartbeat to detect dead connections
				fmt.Fprintf(w, "data: heartbeat\n\n")
				flusher.Flush()

			case <-r.Context().Done():
				// Client disconnected (browser closed, network issue)
				return
			}
		}
	}
}

// HealthCheckHandler - monitor broker health
func HealthCheckHandler(broker *RobustBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats := broker.GetStats()

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "healthy",
			"active_clients": %d,
			"total_clients": %d,
			"messages_sent": %d,
			"messages_dropped": %d,
			"last_message_time": "%s"
		}`,
			stats.ActiveClients,
			stats.TotalClients,
			stats.MessagesSent,
			stats.MessagesDropped,
			stats.LastMessageTime.Format(time.RFC3339))
	}
}

// CleanupHandler - manually trigger cleanup
func CleanupHandler(broker *RobustBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		broker.Cleanup()
		stats := broker.GetStats()

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "cleanup_completed",
			"active_clients": %d
		}`, stats.ActiveClients)
	}
}
