package main

import (
	"context"
	"log"

	"user_authentication/router"

	// "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB URI (replace with your actual URI)
	mongoURI := "mongodb+srv://<username>:<password>.@cluster0.fek5tj1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify the connection
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	// Your application logic here

	router.SetUpRouter(client).Run(":8080")
	// Disconnect the client once your application is done
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}
