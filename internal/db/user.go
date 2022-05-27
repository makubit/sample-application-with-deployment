package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type User struct {
	Username    string    `bson:"username" json:"username"`
	DateOfBirth time.Time `bson:"dateOfBirth" json:"dateOfBirth"`
}

func (conn *MongoConnection) InsertUser(username string, data time.Time) error {
	collection := conn.Cli.Database(conn.Database).Collection(conn.Collection)

	_, err := collection.InsertOne(context.Background(), User{username, data})

	return err
}

func (conn *MongoConnection) GetUser(username string) (User, error) {
	collection := conn.Cli.Database(conn.Database).Collection(conn.Collection)

	var user User
	err := collection.FindOne(context.Background(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
