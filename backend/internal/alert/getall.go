package alert

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(response http.ResponseWriter, request *http.Request, db *mongo.Database) {
	response.Header().Set("Content-Type","application/json")

	collection := db.Collection("Alert")

	var users []Alert

	cursor,err := collection.Find(context.TODO(),bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		// http.Error(w,"Server Error", http.StatusInternalServerError)
		return
	}

	if err := cursor.All(context.TODO(),&users); err != nil {
		response.WriteHeader(http.StatusInternalServerError)

		return
	}

	userEmail := make([]Alert,len(users))
	for i, user := range users {
		userEmail[i] = Alert{Sensor:user.Sensor,Time:user.Time,Alert:user.Alert}
	}


	jsonData, err := json.Marshal(userEmail)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Write(jsonData)
	response.WriteHeader(http.StatusAccepted)
}