package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

const (
	birthDateFormat = "2006-01-02"
	database        = "test"
	collection      = "username"
)

type User struct {
	Username    string    `bson:"username" json:"username"`
	DateOfBirth time.Time `bson:"dateOfBirth" json:"dateOfBirth"`
}

func putUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, newErrorResp(fmt.Errorf("no username provided")))
		return
	}

	var req usernameReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, newErrorResp(err))
		return
	}

	data, err := time.Parse(birthDateFormat, req.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrorResp(err))
		return
	}

	collection := mCli.Database(database).Collection(collection)

	_, err = collection.InsertOne(context.Background(), User{username, data})
	if err != nil {
		log.Fatal(err)
	}

	c.Status(http.StatusNoContent)
}

func getUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, newErrorResp(fmt.Errorf("no username provided")))
		return
	}

	collection := mCli.Database(database).Collection(collection)

	var result User
	err := collection.FindOne(context.Background(), bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrorResp(err))
		return
	}

	msg := createResponse(result)

	c.JSON(http.StatusOK, &usernameResp{
		Message: msg,
	})
}

func isBirthdayToday(date time.Time) bool {
	if time.Now().Month() == date.Month() &&
		time.Now().Day() == date.Day() {
		return true
	}
	return false
}

func calculateBirthday(date time.Time) int {
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	birthdayThisYear := time.Date(time.Now().Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	var days float64
	if birthdayThisYear.Before(today) {
		nextBirthday := birthdayThisYear.AddDate(1, 0, 0)
		days = nextBirthday.Sub(today).Hours() / 24
	} else {
		days = birthdayThisYear.Sub(today).Hours() / 24
	}
	return int(days)
}
