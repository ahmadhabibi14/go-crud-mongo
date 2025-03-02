package database

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	var err error

	mongoUser := os.Getenv(`MONGO_INITDB_ROOT_USERNAME`)
	mongoPass := os.Getenv(`MONGO_INITDB_ROOT_PASSWORD`)
	mongoUrl := fmt.Sprintf("mongodb://%s:%s@localhost:27017", mongoUser, mongoPass)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)

	MongoClient, err = mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
}
