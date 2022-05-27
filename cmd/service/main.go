package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"makubit.com/sample-app/internal/db"
)

var conn *db.MongoConnection

var (
	uri = flag.String("uri", "mongodb://mongodb0.example.com:27017", "connection uri to mongo database")
)

func main() {
	flag.Parse()

	conn = db.NewMongoConnection(*uri)

	e := gin.Default()
	hello := e.Group("/hello")
	{
		hello.PUT("/:username", putUser)
		hello.GET("/:username", getUser)
	}

	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
