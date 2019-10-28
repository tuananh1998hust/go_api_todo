package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient :
func MongoClient() *mongo.Client {
	var dbHost = os.Getenv("DBHOST")
	clientOptions := options.Client().ApplyURI(dbHost)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

// CheckConnection :
func CheckConnection() {
	client := MongoClient()
	// Check Connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB is connected...")
}
