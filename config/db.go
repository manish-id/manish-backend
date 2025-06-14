package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	// Only load .env file if running locally (not on Heroku)
	if os.Getenv("DYNO") == "" {
		// Load .env for local environment
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Get MongoDB URI from environment variables
	MONGOSTRING := os.Getenv("MONGOSTRING")
	if MONGOSTRING == "" {
		log.Fatal("MONGOSTRING is not set")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGOSTRING))
	if err != nil {
		log.Fatal(err)
	}

	// Set the global MongoClient
	MongoClient = client
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}
