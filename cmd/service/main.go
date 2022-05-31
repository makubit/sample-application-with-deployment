package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"makubit.com/sample-app/internal/db"
)

var conn *db.MongoConnection

func main() {
	conn = db.NewMongoConnection()

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
