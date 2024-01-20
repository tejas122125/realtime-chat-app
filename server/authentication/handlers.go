package authentication

import (
	"chatting/server/types"
	"encoding/json"
	"fmt"
	"time"

	// "chatting/server/mongodb"
	"log"

	"github.com/dgrijalva/jwt-go"
	// "go.mongodb.org/mongo-driver/mongo"

	"net/http"
)

//	type Jwt struct {
//		Jwt string `json:"jwt"`
//	}

type ClaimRes struct{
	UserId string `json:"userid"`
}
type UserDetails struct{
	Name string
	email string
}


type UserSigin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}
type Claims struct {
	jwt.StandardClaims
	UserId string `json:"userid"`
}
type errors struct {
	Error string `json:"error"`
}

const SECRET_KEY = "TEJASWEESINGH"
// func keyFunc (token *jwt.Token) (interface{}, error) {
// 	// Your logic to return the key based on the token claims or any other criteria
// 	return []byte("your-secret-key"), nil
// }

func AuthenticationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body := r.Header.Get("jwt")
		// body, err := io.ReadAll(r.Body)
		if body == "" {
			err := errors{Error: "errror while gettin token from user"}
			jsonData, _ := json.Marshal(err)
			log.Println("error wwhile reading from request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonData)
			return
		}
		
		token, err := jwt.ParseWithClaims(
			body,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(SECRET_KEY), nil
			},
		)
		if err != nil {
			err := errors{Error: "error while getting token from user"}
			jsonData, _ := json.Marshal(err)
			log.Println("error wwhile reading from request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonData)
			return

		}
		claims, ok := token.Claims.(*Claims)
		if !ok {
			err := errors{Error: "error while getting token from user"}
			jsonData, _ := json.Marshal(err)
			log.Println("error wwhile reading from request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonData)
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			err := errors{Error: "error while getting token from user"}
			jsonData, _ := json.Marshal(err)
			log.Println("error wwhile reading from request body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonData)
			return
		}
		jsonData, _ := json.Marshal(claims)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)

		fmt.Println("middle ware called")
		next.ServeHTTP(w, r)
	})

}
func HandleCheckSignin(w http.ResponseWriter, r *http.Request) {
	var signin Claims;
	if err:=json.NewDecoder(r.Body).Decode(&signin);err!=nil{
		fmt.Println("error while handleing signin")
		json.NewEncoder(w).Encode(errors{
			Error: "claim not present",
		})
		return
	}
	claimres:= ClaimRes{UserId: signin.Id}
	json.NewEncoder(w).Encode(claimres);
}
func Signin (w http.ResponseWriter,r *http.Request){
var newuser UserSigin;
json.NewDecoder(r.Body).Decode(&newuser)
// data base check up





}
func Signup (w http.ResponseWriter,r *http.Request){
	var userdetails types.UserToDb;
	json.NewDecoder(r.Body).Decode(&userdetails)

	res,err := SaveUserToDatabase(userdetails)
	if err!= nil{
		fmt.Println("error while saving to data base")
		json.NewEncoder(w).Encode(errors{Error: err.Error()})
return

		}
		var mongoresult ClaimRes = ClaimRes{UserId: res.InsertedID.(string)}
		json.NewEncoder(w).Encode(mongoresult)
	}
	// data base login of details
	
