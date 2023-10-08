package authentication

import (
	"encoding/json"

	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
First thing we will be doing is setting the header to return json, so when we use this func in postman, it will return properly styled json.
Then we set user to UserData struct and then read the response passed from frontend by decoding it and setting that to user.
then we make a channel called existUser and genJwt which will be used to call UserExist, and jwt.Generate() having the go in front of the func
means that you are using go routine basically meaning calling a thread to execute this func.
While threads execute the func NewSignUp will make a new uuid and set it to user.Id and then hash the password afterwards by calling HashPassword func
if HashPassword doesn't return any error we will be setting user.Password to hpass, and also set AuthType to Regular since user is making the account manually
**IMPORTANT** to call the thread and stop running this function, look at the way I call userExist := <- existUser, this will wait till the UserExist func is done,
then asssign the value for existUser channel to userExist by using this `<-` and same thing is done with JWT
after getting JWT, we check if user exists or not and if they do exist, we return error 409 and if the user doesn't than we insert the user and
return the JWT
*/
func NewSignUp(response http.ResponseWriter,request *http.Request, db *mongo.Database) {
	
	response.Header().Set("Content-Type","application/json")

	var user UserData
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Mission Json Data"))
		return
	}
	
	user.ID = primitive.NewObjectID()
	

	existUser := make(chan string)

	go UserExist(user,db,existUser)

	hpass,err := HashPassword(user.Password)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error hashing password"))
		return
	}
	user.Password = string(hpass)
	user.AuthType = "Regular"
	userExist := <- existUser
	if userExist != "" {
		response.WriteHeader(http.StatusConflict)
		response.Write([]byte("User Already Exist"))
		return
	} else {
		err := InsertUser(user,db)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("Query Error"))
			return
		}
	}
	response.WriteHeader(http.StatusCreated)
}


