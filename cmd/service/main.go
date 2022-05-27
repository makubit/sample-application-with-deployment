package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"makubit.com/sample-app/internal/db"
)

var mCli *mongo.Client

var (
	uri = flag.String("uri", "mongodb://mongodb0.example.com:27017", "connection uri to mongo database")
)

func main() {
	flag.Parse()

	mCli = db.NewMongoClient(*uri)

	e := gin.Default()
	hello := e.Group("/hello")
	{
		hello.PUT("/:username", putUsername)
		hello.GET("/:username", getUsername)
	}

	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
