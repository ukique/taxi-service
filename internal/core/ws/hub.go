package ws

type Hub struct {
	clients          map[*Client]bool
	register         chan *Client
	unregister       chan *Client
	broadcastOrders  chan []byte
	broadcastDrivers chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:          make(map[*Client]bool),
		register:         make(chan *Client),
		unregister:       make(chan *Client),
		broadcastOrders:  make(chan []byte),
		broadcastDrivers: make(chan []byte),
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
		case orders := <-h.broadcastOrders:
			for client := range h.clients {
				if client.subscribeType == "subscribe_orders" {
					select {
					case client.send <- orders:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		case drivers := <-h.broadcastDrivers:
			for client := range h.clients {
				if client.subscribeType == "subscribe_drivers" {
					select {
					case client.send <- drivers:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}

}
