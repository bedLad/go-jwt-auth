package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB!!")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, c_name string) *mongo.Collection {
	collection := client.Database("golang-auth-db").Collection(c_name)
	return collection
}
