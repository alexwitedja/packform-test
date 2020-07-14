package mongohelper

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB : This is helper function to connect mongoDB
func ConnectDB() *mongo.Database {
	// Change mongodb uri so it matches your local machine.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	// if err is not nil, then client will be nil, and applications should not Disconnect a nil client
	if err != nil {
		log.Fatal(err)
	}

	// Connect doesn't actually error if any (or all) of the servers are unreachable
	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to MongoDB!")

	db := client.Database("packformtest")

	return db
}
