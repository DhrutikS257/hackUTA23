package alert

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Get(response http.ResponseWriter,request *http.Request, db *mongo.Database) {
	collection := db.Collection("User")
	var alert Alert
	var email Email
	err := json.NewDecoder(request.Body).Decode(&alert)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error Decoding Json"))
		return
	}
	sendEmail := make(chan bool)
	insert := make(chan error)

	err = collection.FindOne(context.Background(),alert.Email).Decode(&email)
	if err == mongo.ErrNoDocuments{
		response.WriteHeader(http.StatusNotFound)
		return
	}
	go Send(alert.Email,alert.Alert,sendEmail)
	go Insert(alert,db,insert)


	emailSent := <- sendEmail
	dataInserted := <- insert

	if emailSent && dataInserted == nil {
		response.WriteHeader(http.StatusAccepted)
		return
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

}