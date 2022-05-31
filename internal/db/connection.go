package db

import (
	"context"
	"log"
	"os"
	"time"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	secretPath = "/etc/mongodb/auth/"
	secretName = "uri"
	database   = "test"
	collection = "username"
)

type MongoConnection struct {
	Cli        *mongo.Client
	Database   string
	Collection string
}

func NewMongoConnection() *MongoConnection {
	uri := retrieveUri()
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConnection{
		Cli:        client,
		Database:   database,
		Collection: collection,
	}
}

func retrieveUri() string {
	uri, err := os.ReadFile(secretPath + secretName)
	if err != nil {
		return "mongodb://localhost:27017"
	}
	return string(uri)
}
