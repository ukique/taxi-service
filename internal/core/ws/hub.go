package ws

import (
	"encoding/json"
	"log"

	"github.com/ukique/taxi-service/internal/models"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) SendToBroadcast(payload []byte) {
	select {
	case h.broadcast <- payload:
	default:
		log.Println("high-load: ws broadcast channel is full!")
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			var messageBody models.OutgoingMessage[any]
			if err := json.Unmarshal(message, &messageBody); err != nil {
				log.Println("unresolved ws message in broadcast:", err)
				log.Println("message:", string(message))
				continue
			}
			for client := range h.clients {
				if client.subscribeType == messageBody.Type {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}

}
