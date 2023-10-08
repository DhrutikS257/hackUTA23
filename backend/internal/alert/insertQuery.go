package alert

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(alert Alert,db *mongo.Database,ch chan error) {
	collection := db.Collection("Alert")

	_,err := collection.InsertOne(context.Background(),alert)

	ch <- err
}