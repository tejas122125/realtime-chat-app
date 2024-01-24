package types



type Userfromdb struct{
	
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}
type UserToDb struct {
	Name string  `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}