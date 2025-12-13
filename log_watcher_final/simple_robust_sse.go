package main

import (
	"fmt"
	"net/http"
)

// SimpleRobustSSEHandler - simple SSE with basic error handling
func SimpleRobustSSEHandler(broker *SimpleRobustBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create client channel with small buffer
		clientChan := make(chan string, 10)

		// Register client
		broker.AddClient(clientChan)

		// Clean up when done
		defer broker.RemoveClient(clientChan)

		// Send messages to browser
		for {
			select {
			case message := <-clientChan:
				fmt.Fprintf(w, "data: %s\n\n", message)
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			case <-r.Context().Done():
				// Client disconnected
				return
			}
		}
	}
}
