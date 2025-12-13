package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create simple robust broker
	broker := NewSimpleRobustBroker()

	// Start log tailing
	logFilePath := "./app.log"
	go SimpleRobustTailFile(logFilePath, broker)

	// Set up HTTP routes
	http.HandleFunc("/events", SimpleRobustSSEHandler(broker))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Simple Robust Log Watcher Server started at :8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
