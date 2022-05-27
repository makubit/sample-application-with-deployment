package main

import (
	"fmt"
	"makubit.com/sample-app/internal/db"
)

type usernameReq struct {
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
}

type usernameResp struct {
	Message string `json:"message" binding:"required"`
}

type errorResp struct {
	Message string `json:"message"`
}

func newErrorResp(err error) *errorResp {
	return &errorResp{
		Message: err.Error(),
	}
}

func createResponse(user db.User) string {
	var days int
	if !isBirthdayToday(user.DateOfBirth) {
		days = calculateBirthday(user.DateOfBirth)
		return fmt.Sprintf("Hello, %s! Your birthday is in %d day(s)", user.Username, days)
	}
	return fmt.Sprintf("Hello, %s! Happy birthday!", user.Username)
}
