package alert

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteAll(response http.ResponseWriter,request http.Request,db *mongo.Database) {
	collection := db.Collection("Alert")

	_,err := collection.DeleteMany(context.Background(),bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusAccepted)
}