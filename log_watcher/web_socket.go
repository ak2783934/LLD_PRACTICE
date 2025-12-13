package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func readAndWrite() {
	for {
		messageType, p, err := WebSocketConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		newMessage := string(p) + " response"
		bytes := []byte(newMessage)

		if err := WebSocketConn.WriteMessage(messageType, bytes); err != nil {
			log.Println(err)
			return
		}
	}
}

func PublishMessage(message string) {
	fmt.Println("message published")
	fmt.Println(message)
	if WebSocketConn == nil {
		fmt.Println("web socket connection not found")
		return
	}
	if err := WebSocketConn.WriteMessage(1, []byte(message)); err != nil {
		log.Println(err)
		return
	}
}
