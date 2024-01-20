package main

import (
	"chatting/server/authentication"
	"chatting/server/ws"
	"chatting/server/appwrite"

	// "os"

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
	// router.HandleFunc("/check",func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("checked")

	// })
	authenticate := router.PathPrefix("/auth").Subrouter()
	r := router.PathPrefix("/user").Subrouter()

// authentication routes
 authenticate.Use(authentication.AuthenticationMiddleware)
authenticate.HandleFunc("/checksignin",authentication.HandleCheckSignin).Methods("POST")
r.HandleFunc("/signin",authentication.Signin)


// r.HandleFunc("/redirect",func(w http.ResponseWriter, r *http.Request) {
// 	// json.NewEncoder(w).Encode(message{name: "tejaswee",lastname: "singh"})
// 	http.Redirect(w,r,"https://www.example.com/",http.StatusAccepted)
// })
// // var msg message;
// r.HandleFunc("/heredirected",func(w http.ResponseWriter, r *http.Request) {
// 	// json.NewDecoder(r.Body).Decode(&msg)
// 	// fmt.Println(msg)
// 	fmt.Println("called here")
// })
r.HandleFunc("/ws/joinroom",wshandler.JoinRoom)
r.HandleFunc("/ws/getrooms",wshandler.GetRooms)
r.HandleFunc("/ws/getclients",wshandler.GetClients)
r.HandleFunc("/ws/createroom",wshandler.CreateRoom)
	if err:=http.ListenAndServe(":8080",router);err!=nil{
		fmt.Println("there is a error in init routers function ")
	}


}

func main() {
	// get:=os.Getenv("APPWRITE_PROJECT_ID")
	fmt.Println("get")
appwrite.Setup()
	// h := ws.NewHub()
	// wshandler := ws.NewHandler(h)
	// go h.Run()
	// initRouters(wshandler)

}