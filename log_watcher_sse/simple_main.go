package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create simple broker
	broker := NewSimpleBroker()

	// Start log tailing in background
	logFilePath := "./app.log"
	go SimpleTailFile(logFilePath, broker)

	// Set up HTTP routes
	http.HandleFunc("/events", SimpleSSEHandler(broker))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Simple Log Watcher Server started at :8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
