package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URL from .env
	MongoDb := os.Getenv("MONGODB_URL")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// Initialize the client as a global variable
var Client *mongo.Client = DBinstance()

// Function to open a specific collection in the PortfoAI database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Set database name to "PortfoAI"
	var collection *mongo.Collection = client.Database("PortfoAI").Collection(collectionName)
	return collection
}
