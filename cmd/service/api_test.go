package main

import (
	"fmt"
	"makubit.com/sample-app/internal/db"
	"testing"
	"time"
)

func TestNewErrorResponse(t *testing.T) {
	var tests = []struct {
		err  error
		want errorResp
	}{
		{fmt.Errorf("Some Error"), errorResp{Message: "Some Error"}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.err)
		t.Run(testname, func(t *testing.T) {
			res := newErrorResp(tt.err)
			if res.Message != tt.want.Message {
				t.Errorf("got %v, want %v", res, tt.want)
			}
		})
	}
}

func TestCreateResponse(t *testing.T) {
	var tests = []struct {
		user db.User
		want string
	}{
		{db.User{
			Username:    "username1",
			DateOfBirth: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
		}, "Hello, username1! Happy birthday!"},
		{db.User{
			Username:    "username2",
			DateOfBirth: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.UTC),
		}, "Hello, username2! Your birthday is in 1 day(s)"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.user.Username)
		t.Run(testname, func(t *testing.T) {
			res := createResponse(tt.user)
			if res != tt.want {
				t.Errorf("got %v, want %v", res, tt.want)
			}
		})
	}
}
