package authentication

import (
	
	"chatting/server/types"
	// "errors"

	"context"
	"fmt"
	"time"
// "chatting/server/authentication"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserById(userid string)(*mongo.SingleResult) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := GetDatabase()
	collection := OpenCollection(client, "user")
	user := collection.FindOne(ctx,bson.D{{Key: "_id",Value: userid}})
	return user
	// decode if user !== mongo.errnodocuments


}
func GetCurrentUser() {

}
func SaveUserToDatabase(userrdetails types.UserToDb) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := GetDatabase()
	collection := OpenCollection(client, "user")
	res, err := collection.InsertOne(ctx, userrdetails)
	if err != nil {
		fmt.Println("error while saving user to database")
		return res, err
	}
	return res, nil
}

func GetUserByEmail(user UserSigin) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := GetDatabase()
	collection := OpenCollection(client, "user")
	signeduser := collection.FindOne(ctx,bson.D{{Key: "email",Value: user.Email}})
	return signeduser
}

