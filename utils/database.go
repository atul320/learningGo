package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(os.Getenv("DB_NAME"))
	log.Println("Connected to MongoDB!")
}

func CloseDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := DB.Client().Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB!")
}
