package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserFromDb struct {
	Id  primitive.ObjectID `bson:"_id"`
	Name string  `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
}
type UserToDb struct {
	Name string  `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}