package db

import (
	"context"
	"log"
	"time"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Cli        *mongo.Client
	Database   string
	Collection string
}

func NewMongoConnection(uri string) *MongoConnection {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConnection{
		Cli:        client,
		Database:   "test",
		Collection: "username",
	}
}
