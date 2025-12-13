package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "server running!\n")
}

var WebSocketConn *websocket.Conn

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	WebSocketConn = conn
	fmt.Println("web socket connection created")
	// publish the last 10 lines here somehow?
	// readAndWrite()
}

func StartServer() {
	http.HandleFunc("/health-check", healthCheck)
	http.HandleFunc("/ws", wsEndpoint)
	fmt.Println("server started")
	go startConsumerForTextFile()
	http.ListenAndServe(":3000", nil)
}
