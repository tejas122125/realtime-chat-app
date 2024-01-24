package main

import (
	"chatting/server/authentication"
	"chatting/server/ws"
	// "chatting/server/authentication"
    "github.com/joho/godotenv"
	"os"

	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type message struct{
	name string `json : "name"`
	lastname string `json : "lastname"`
}
func initRouters(wshandler *ws.Handler) {
	router  :=  mux.NewRouter()
	
	authenticate := router.PathPrefix("/auth").Subrouter()
	r := router.PathPrefix("/user").Subrouter()

// authentication routes
 authenticate.Use(authentication.AuthenticationMiddleware)
authenticate.HandleFunc("/checksignin",authentication.HandleCheckSignin).Methods("POST")



r.HandleFunc("/signin",authentication.Signin).Methods("POST")
r.HandleFunc("/signup",authentication.Signup).Methods("POST")
r.HandleFunc("/ws/joinroom",wshandler.JoinRoom)
r.HandleFunc("/ws/getrooms",wshandler.GetRooms)
r.HandleFunc("/ws/getclients",wshandler.GetClients)
r.HandleFunc("/ws/createroom",wshandler.CreateRoom)
	if err:=http.ListenAndServe(":8080",router);err!=nil{
		fmt.Println("there is a error in init routers function ")
	}


}

func main() {
	err := godotenv.Load()
	if err != nil {
        fmt.Println("Error loading .env file")
        return
    }
	get:=os.Getenv("MONGO_URL")
	fmt.Println(get,"vgdcgfcgfdcgf")
// authentication.Setup()
	h := ws.NewHub()
	wshandler := ws.NewHandler(h)
	go h.Run()
	initRouters(wshandler)

}