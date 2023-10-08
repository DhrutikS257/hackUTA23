package authentication

import (
	"context"


	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

/*
This query will insert the user in the database, we are passing UserData which will have ID, Name, Email, Password, and AuthType
and second parameter is the connection pool. If there was an error inserting user, it will be handled in signup.go and callback.go
*/
func InsertUser(user UserData, db *mongo.Database) error{
	collection := db.Collection("User")
	_,err := collection.InsertOne(context.Background(),user)

	return err
}

/*
This query will check if the user exists, we are passing UserData and connection pool, along with that we are passing channel which will
have boolean value, ch is used to call this function concurrently, after doing the query we will check if err == NoRows meaning we didn't get
anything from the query, and assign that boolean value to channel
*/
func UserExist(user UserData, db *mongo.Database, ch chan string){
	collection := db.Collection("User")
	filter := bson.M{"email":user.Email}
	var queryUser UserData
	err := collection.FindOne(context.Background(),filter).Decode(&queryUser)
	if err == mongo.ErrNoDocuments || err != nil {
		ch <- ""
	} else {
		ch <- queryUser.ID.Hex()
	}
}

func GetUser(user UserData, db *mongo.Database) string{
	collection := db.Collection("User")

	filter := bson.M{"email":user.Email}
	var queryUser UserData

	err := collection.FindOne(context.Background(),filter).Decode(&queryUser)
	if err == mongo.ErrNoDocuments || err != nil {
		return ""
	}
	return queryUser.Password
}


/*
We are hasing the user password, here by passing password as a parameter, and it will return the hash password along with error if there is any. 
We are using bcrypt for hashing password
*/
func HashPassword(password string) ([]byte,error){
	hpass,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return hpass,err
}

func VerifyPassword(hash []byte, password []byte) bool{
	err := bcrypt.CompareHashAndPassword(hash,password)
	return err == nil
}


