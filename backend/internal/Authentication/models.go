package authentication

import "go.mongodb.org/mongo-driver/bson/primitive"



type UserData struct {
	ID primitive.ObjectID `bson:"_id"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	AuthType string `bson:"authtype"`
}

type LoginData struct {
	Email string `bson:"email"`
	Password string `bson:"password"`
}