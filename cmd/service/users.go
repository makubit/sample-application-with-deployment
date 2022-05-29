package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	birthDateFormat = "2006-01-02"
)

func putUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusNotFound, newErrorResp(fmt.Errorf("no username provided")))
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

	err = conn.InsertUser(username, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrorResp(err))
		return
	}

	c.Status(http.StatusNoContent)
}

func getUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusNotFound, newErrorResp(fmt.Errorf("no username provided")))
		return
	}

	user, err := conn.GetUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrorResp(err))
		return
	}

	c.JSON(http.StatusOK, &usernameResp{
		Message: createResponse(user),
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
