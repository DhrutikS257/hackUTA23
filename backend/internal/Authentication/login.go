package authentication

import (
	"backend/internal/alert"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Login(response http.ResponseWriter, request *http.Request,db *mongo.Database) {
	response.Header().Set("Content-Type","application/json")

	var user UserData
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error Decoding Json"))
		return
	}
	sendEmail := make(chan bool)
	go alert.Send(user.Email,"THREAT",sendEmail)

	// checkPassword := make(chan bool)
	password := GetUser(user,db)
	if  password == ""{
		response.WriteHeader(http.StatusNotFound)
		return
	}
	matchPassword := VerifyPassword([]byte(password),[]byte(user.Password))
	emailSent := <- sendEmail
	if emailSent {
		if matchPassword {
			response.WriteHeader(http.StatusAccepted)
			return
		} else {
			response.WriteHeader(http.StatusNotFound)
			return
		}
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}