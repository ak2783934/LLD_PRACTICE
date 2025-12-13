package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// TestSimpleLogWatcher - Test the simple version with last 10 lines
func TestSimpleLogWatcher() {
	// Create simple broker
	broker := NewSimpleBroker()

	// Simulate some log lines being added to the buffer
	fmt.Println("Adding some test log lines...")
	for i := 1; i <= 15; i++ {
		line := fmt.Sprintf("Log line %d - %s", i, time.Now().Format("15:04:05"))
		broker.Broadcast(line)
		time.Sleep(100 * time.Millisecond) // Small delay to see the effect
	}

	// Start log tailing in background
	logFilePath := "./app.log"
	go SimpleTailFile(logFilePath, broker)

	// Set up HTTP routes
	http.HandleFunc("/events", SimpleSSEHandler(broker))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Simple Log Watcher Server started at :8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	fmt.Println("You should see the last 10 log lines when you connect!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
