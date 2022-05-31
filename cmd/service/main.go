package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"makubit.com/sample-app/internal/db"
)

var conn *db.MongoConnection

func main() {
	setupMongoDBConnection()

	e := setupServer()
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}

func setupServer() *gin.Engine {
	e := gin.Default()
	hello := e.Group("/hello")
	{
		hello.PUT("/:username", putUser)
		hello.GET("/:username", getUser)
	}
	return e
}

func setupMongoDBConnection() {
	conn = db.NewMongoConnection()
}
