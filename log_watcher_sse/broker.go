package main

type Broker struct {
	Clients  map[chan string]bool // Connected clients
	Add      chan chan string     // New client
	Remove   chan chan string     // Disconnected client
	Messages chan string          // New messages to broadcast
}

func NewBroker() *Broker {
	return &Broker{
		Clients:  make(map[chan string]bool),
		Add:      make(chan chan string),
		Remove:   make(chan chan string),
		Messages: make(chan string),
	}
}

func (broker *Broker) Start() {
	for {
		select {
		case client := <-broker.Add:
			broker.Clients[client] = true
		case client := <-broker.Remove:
			delete(broker.Clients, client)
			close(client)
		case msg := <-broker.Messages:
			for client := range broker.Clients {
				client <- msg
			}
		}
	}
}
