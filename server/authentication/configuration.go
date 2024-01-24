package authentication

// import (
// 	"fmt"
// 	// "os"
// 	// "github.com/appwrite/sdk-for-go/database"
// 	"github.com/appwrite/sdk-for-go"
// )
// type usertest struct {
// 	name string
// 	email string
// password string
// userid string
// }
// // type AppwriteConfig struct{
// // 	Client appwrite.Client `json:"client"`
// // 	Database appwrite.Database `json:"database"`
// // }

// // to be called inittially toi set up configuration value

// func Setup (){
// 	var clients = appwrite.NewClient()
// 	// clients.SetProject(os.Getenv("APPWRITE_PROJECT_ID"))
// 	// clients.SetKey(os.Getenv("APPWRITE_SECRET_API_KEY"))
// 	clients.SetProject("64adaa390de34d88cefe")
// 	clients.SetKey("ee18c41a054e7a1d965f26c23fb0c2bbd79c55ed5563ca7fd2a1f2b66bbd1d52c16d43bd4accb9d2022de15448b8b97e0585fa92182436ec531793a0e5b3b25cfd6d654906ece2676e76b1104e72f58a7eaded88b2721963335edb20a1cae8689e0bdde3afbd45902d7d82874ea608ee5cf1b9c583d7779efe4ee6fb7e9c276d")
// 	clients.SetEndpoint("https://cloud.appwrite.io/v1")
// 	var ser = appwrite.NewDatabase(clients)

//	// var service = appwrite.Database{clients: &clients}
//
// fmt.Println("heree reacgersdeds")
//
//		document,err := ser.CreateDocument("users1234",usertest{name: "tejas",email: "tejas",password: "teljfdkl",userid: "tererer"},nil,nil,"","","")
//		if err != nil {
//			fmt.Println("error while creating document")
//		}
//		fmt.Println("documentis",document)
//	}

import (
	"context"
	"fmt"
	"log"


	// "os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("error while loading .env file in go")
	}
	// mongoDb := os.Getenv("MONGODB_URL")
	const mongoDb = "mongodb+srv://tejasweekumarsingh:o8rftH6R1GMs0YW1@tejas.tokgflw.mongodb.net"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Println("mongo error",err)
		return nil
	}
	fmt.Println("mongo connection successful")
	return client

}
func OpenCollection (client *mongo.Client,collectionName string )*mongo.Collection{

fmt.Println("collection goot successfully")
	var collection *mongo.Collection = client.Database("chatapp").Collection(collectionName)
	return collection
}
