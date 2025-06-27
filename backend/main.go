package main

import (
	"fmt"

	"fixit.com/backend/internal"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Mongo driver
const MONGO_URL = "mongodb://localhost:27017"
const MONGO_DB = "fixit"

func createMongoClient() (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	return mongoClient, nil
}

func main() {
	mongoClient, err := createMongoClient()
	if err != nil {
		panic(err)
	}

	err = internal.StartHttpServer(mongoClient)
	if err != nil {
		panic(err)
	}
}
