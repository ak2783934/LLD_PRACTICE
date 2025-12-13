package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketServer struct {
	logTail *LogTail
}

func NewWebSocketServer(logTail *LogTail) *WebSocketServer {
	return &WebSocketServer{logTail: logTail}
}

func (ws *WebSocketServer) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Add subscriber to log tail
	subscriberID, ch := ws.logTail.AddSubscriber()
	defer ws.logTail.RemoveSubscriber(subscriberID)

	// Send last 10 lines immediately
	for _, line := range ws.logTail.GetLastLines() {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
			return
		}
	}

	// Stream new lines
	for {
		select {
		case line := <-ch:
			if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
				return
			}
		}
	}
}
