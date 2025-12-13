package main

import (
	"fmt"
	"net/http"
)

// SimpleSSEHandler - much simpler SSE handler
func SimpleSSEHandler(broker *SimpleBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create a channel for this client
		clientChan := make(chan string)

		// Register client with broker
		broker.AddClient(clientChan)

		// Clean up when client disconnects
		defer broker.RemoveClient(clientChan)

		// Stream messages to client
		for {
			select {
			case message := <-clientChan:
				// Send message to browser in SSE format
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
