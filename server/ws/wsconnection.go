package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {

	var req CreateRoomReq
	// fmt.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	h.hub.Rooms[req.ID] = &Room{ID: req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(req)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

}

func (h *Handler) JoinRoom (w http.ResponseWriter,r * http.Request){
	// upgrader.CheckOrigin = func(r *http.Request) bool {return true}
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}
// w.Header().Set("upgrade",upgrader)

	conn,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		fmt.Println("error while joining the room",err)
		http.Error(w,err.Error(),http.StatusBadGateway)
	}

	params := r.URL.Query()
	id := params.Get("id")
	clientid:=params.Get("clientid")
	username:= params.Get("username")
	fmt.Println(id +" "+clientid+" "+username+" ")
	cl := &Client{
		Conn : conn,
		Message: make(chan *Message),
		ID:clientid ,
		RoomID:id ,
		Username: username,
	}
	msg := &Message{
		Content: `new user joined of id ${clientid}`,
		RoomID: id,
		Username: username,
	}
	h.hub.Register<-cl;
	h.hub.Broadcast<-msg;
	go cl.Write();
	go cl.Read(h.hub)
}
