package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	broker := NewBroker()
	go broker.Start()

	logFilePath := "./app.log" // path to your log file
	go TailFile(logFilePath, broker)

	http.HandleFunc("/events", SSEHandler(broker))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
