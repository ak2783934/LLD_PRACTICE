package main

import (
	"log"
	"net/http"
)

func main() {
	logFile := "/tmp/dat1" // Replace with your log file path

	logTail := NewLogTail(logFile)
	logTail.Start()

	wsServer := NewWebSocketServer(logTail)

	http.HandleFunc("/ws", wsServer.HandleWS)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs) // frontend files

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
