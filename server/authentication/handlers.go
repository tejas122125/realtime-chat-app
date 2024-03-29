package authentication

import (
	"chatting/server/types"
	"encoding/json"
	"errors"
	"fmt"

	// "go/token"
	"time"

	// "chatting/server/mongodb"
	// "log"
	// "github.com/joho/godotenv"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	type Jwt struct {
//		Jwt string `json:"jwt"`
//	}

type ClaimRes struct {
	UserId string `json:"userid"`
}
type UserDetails struct {
	Name  string
	email string
}

type UserSigin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ClaimsJwt struct {
	jwt.StandardClaims
	UserId string `json:"userid"`
}
type cuserrors struct {
	Error string `json:"error"`
}

// err := godotenv.Load(".env")
const SECRET_KEY = "tejaswee"

func generatejwttoken(userid string)(string,error){
	newclaim := &ClaimsJwt{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Local().Add(time.Hour*60).Unix()),
		},
	}

	fmt.Println(newclaim)

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256,newclaim).SignedString([]byte(SECRET_KEY))
	if err != nil {
		fmt.Println("error while signing the jwt token")
		fmt.Println(err)
		return token , err
	}
	return token,nil

}

func Validatetoken (tokenrecieved string)(*ClaimsJwt,error){
	
	token, err := jwt.ParseWithClaims(
		tokenrecieved,
		&ClaimsJwt{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)	
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsJwt)
	if !ok {
		fmt.Println("error in getting jwt")
		return nil,errors.New("error in getting jwtr")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		fmt.Println("error in getting jwt")
		return nil,errors.New("jwt expired")
	}

return claims,nil

}



func AuthenticationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			fmt.Fprintln(w, "Cookie not found")
			return
		}
	
		// Access the cookie value
		jwt_token := cookie.Value

		// body := r.Header.Get("jwt")
		w.Header().Set("Content-Type", "application/json")
		token,err :=  Validatetoken(jwt_token)
		if err != nil{
			json.NewEncoder(w).Encode(errors.New(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(token)
		fmt.Println("middle ware called")
		next.ServeHTTP(w, r)
	})

}
func HandleCheckSignin(w http.ResponseWriter, r *http.Request) {
	var signin ClaimsJwt
	if err := json.NewDecoder(r.Body).Decode(&signin); err != nil {
		fmt.Println("error while handleing signin")
		json.NewEncoder(w).Encode(errors.New("error while handleing signin"))
		return
	}
	claimres := ClaimRes{UserId: signin.Id}
	json.NewEncoder(w).Encode(claimres)
}
func Signin(w http.ResponseWriter, r *http.Request) {
	var newuser UserSigin
	json.NewDecoder(r.Body).Decode(&newuser)
	// data base check up
	fmt.Println(newuser.Email)
	fmt.Println(r.Body)

	res := GetUserByEmail(newuser)

	if res == nil {
		fmt.Println("user not found")
		json.NewEncoder(w).Encode(errors.New("user not found in  documents"))
	}
	fmt.Println(res)
	// fmt.Fprintln(w,res)
	json.NewEncoder(w).Encode(res)

}
func Signup(w http.ResponseWriter, r *http.Request) {
// generating new jwt token

	var userdetails types.UserToDb
	json.NewDecoder(r.Body).Decode(&userdetails)
    fmt.Println(userdetails)
	res, err := SaveUserToDatabase(userdetails)
	if err != nil {
		fmt.Println("error while saving to data base")
		json.NewEncoder(w).Encode(errors.New( err.Error()))
		return

	}
	insertedID := res.InsertedID.(primitive.ObjectID)

	// Convert ObjectID to string
	insertedIDString := insertedID.Hex()
	fmt.Println(insertedIDString,"rtjggjgjhre")
	newtoken,err := generatejwttoken(insertedIDString)
    

	if err != nil {
		fmt.Println("error while getting neew token")
		json.NewEncoder(w).Encode(errors.New( err.Error()))
		return
	}

	cookie := http.Cookie{
		Name:    "jwt",
		Value:   newtoken,
		Expires: time.Now().Add(60 * time.Hour), // Cookie expiration time (adjust as needed)
	}

	http.SetCookie(w, &cookie)
	var mongoresult ClaimRes = ClaimRes{UserId: insertedIDString}
	json.NewEncoder(w).Encode(mongoresult)
}

