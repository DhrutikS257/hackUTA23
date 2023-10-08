package alert

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alert struct {
	Sensor string `bson:"sensor"`
	Time string `bson:"time"`
	Alert string `bson:"alert"`
	Email string `bson:"email"`
}

type Email struct {
	ID primitive.ObjectID `bson:"_id"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	AuthType string `bson:"authtype"`
}