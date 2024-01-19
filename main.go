package main

import (
	"chatting/server/ws"
	"fmt"
	"net/http"
)
func initRouters(wshandler *ws.Handler) {

	http.HandleFunc("/ws/createroom",wshandler.CreateRoom)
	http.HandleFunc("/ws/joinroom",wshandler.JoinRoom)
	http.HandleFunc("/ws/getrooms",wshandler.GetRooms)
	http.HandleFunc("/ws/getclients",wshandler.GetClients)
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		fmt.Println("there is a error in init routers function ")
	}


}

func main() {

	h := ws.NewHub()
	wshandler := ws.NewHandler(h)
	go h.Run()
	initRouters(wshandler)

}