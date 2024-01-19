package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	// "golang.org/x/tools/go/analysis/passes/httpresponse"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}
type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func (cl *Client) Write() {
	defer func() {
		cl.Conn.Close()
	}()
	for {
		message, ok := <-cl.Message
		if !ok {
			return
		}
		cl.Conn.WriteJSON(message)
	}

}
func (cl *Client) Read(hub *Hub) {

	defer func() {
		hub.Unregister <- cl
		cl.Conn.Close()
	}()

	for {

		_, m, err := cl.Conn.ReadMessage()
		if err != nil {
			fmt.Println("error in reading message from client", err)
		}
		hub.Broadcast <- &Message{Content: string(m),
			RoomID:   cl.RoomID,
			Username: cl.Username,
		}

	}
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]RoomRes, 0)

	for _, rm := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   rm.ID,
			Name: rm.Name,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)

}
type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients (w http.ResponseWriter, r *http.Request){
	var clients []ClientRes;
roomId := r.URL.Query().Get("roomid");


if _, ok := h.hub.Rooms[roomId]; !ok {
	clients = make([]ClientRes, 0)
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(responseData)
	
}
for _, c := range h.hub.Rooms[roomId].Clients {
	clients = append(clients, ClientRes{
		ID:       c.ID,
		Username: c.Username,
	})
}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)

}