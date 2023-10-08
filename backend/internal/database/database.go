package database

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
This func will start the connection to cockroachdb using pgx connection pool
what we are doing is getting the connection key from env and setting it to conn_str
then we want to parseconfig which basically means that set the value that are needed for connection pooling
for parseconfig, it will have values like minimum/maximum connection which will explicitely set those but we are just using default values which is good enough
then after setting up parseconfig, we call connectconfig which actually establishes the connection to the database and makes the connection pool which is then returned
most of the time when you call the func it usually returns the value you requested and and err, you always have to check err to see if we got an err, and if we did
display the error
*/


func Connect() *mongo.Database{

	mongo_uri := os.Getenv("MONGO_CONN")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongo_uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	db := client.Database("HackUTA23")

	return db
}