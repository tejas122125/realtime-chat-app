package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (hub *Hub) Run() {

	for {

		select {
		case client := <-hub.Register:
			if _, ok := hub.Rooms[client.RoomID]; ok {
				r := hub.Rooms[client.RoomID]
				if _, ok := r.Clients[client.ID]; !ok {
					r.Clients[client.ID] = client
				}
			}

		case client := <-hub.Unregister:
			if _, ok := hub.Rooms[client.RoomID]; ok {
				if _, ok := hub.Rooms[client.ID].Clients[client.ID]; ok {
					if len(hub.Rooms[client.RoomID].Clients) > 0 {
						hub.Broadcast <- &Message{Content: `user of id ${client.ID} got disconnected`,
							RoomID:   client.RoomID,
							Username: client.Username}
					}
					delete(hub.Rooms[client.ID].Clients, client.ID)
					close(client.Message)

				}

			}

		case message := <-hub.Broadcast:
			if _, ok := hub.Rooms[message.RoomID]; ok {
				for _, cl := range hub.Rooms[message.RoomID].Clients {
					cl.Message <- message
				}
			}
		}
	}
}
