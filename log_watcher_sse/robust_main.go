package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create robust broker
	broker := NewRobustBroker()

	// Start log tailing with rotation handling
	logFilePath := "./app.log"
	go TailFileWithRotation(logFilePath, broker)

	// Start cleanup routine
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				broker.Cleanup()
				stats := broker.GetStats()
				log.Printf("Broker stats - Active: %d, Total: %d, Sent: %d, Dropped: %d",
					stats.ActiveClients, stats.TotalClients, stats.MessagesSent, stats.MessagesDropped)
			}
		}
	}()

	// Set up HTTP routes
	http.HandleFunc("/events", RobustSSEHandler(broker))
	http.HandleFunc("/health", HealthCheckHandler(broker))
	http.HandleFunc("/cleanup", CleanupHandler(broker))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Graceful shutdown handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down gracefully...")
		os.Exit(0)
	}()

	fmt.Println("Robust Log Watcher Server started at :8080")
	fmt.Println("Endpoints:")
	fmt.Println("  /events  - SSE stream")
	fmt.Println("  /health  - Health check")
	fmt.Println("  /cleanup - Manual cleanup")
	fmt.Println("Open http://localhost:8080 in your browser")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
