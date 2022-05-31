package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIsBirthdayToday(t *testing.T) {
	var tests = []struct {
		date time.Time
		want bool
	}{
		{time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC), true},
		{time.Date(time.Now().Year(), time.Now().Month()+1, time.Now().Day(), 0, 0, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.date)
		t.Run(testname, func(t *testing.T) {
			res := isBirthdayToday(tt.date)
			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}

func TestCalculateBirthday(t *testing.T) {
	daysInYear := time.Date(time.Now().Year(), time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	var tests = []struct {
		date time.Time
		want int
	}{
		{time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, time.UTC), daysInYear - 1},
		{time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.UTC), 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.date)
		t.Run(testname, func(t *testing.T) {
			res := calculateBirthday(tt.date)
			if res != tt.want {
				t.Errorf("got %d, want %d", res, tt.want)
			}
		})
	}
}

func TestPutUser(t *testing.T) {
	today := time.Date(time.Now().Year()-20, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Format(birthDateFormat)
	tomorrow := time.Date(time.Now().Year()-20, time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.UTC).Format(birthDateFormat)

	tests := []struct {
		name         string
		expectedCode int
		body         string
		username     string
	}{
		{
			name:         "Status No Content Today",
			expectedCode: http.StatusNoContent,
			body:         fmt.Sprintf("{\"dateOfBirth\":\"%s\"}", today),
			username:     "username1",
		},
		{
			name:         "Status No Content Tomorrow",
			expectedCode: http.StatusNoContent,
			body:         fmt.Sprintf("{\"dateOfBirth\":\"%s\"}", tomorrow),
			username:     "username2",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			setupMongoDBConnection()
			e := setupServer()
			w := performRequest(e, "PUT", fmt.Sprintf("/hello/%s", tc.username), strings.NewReader(tc.body))
			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name         string
		expectedCode int
		expectedBody string
		username     string
	}{
		{
			name:         "Not Found",
			expectedCode: http.StatusNotFound,
			expectedBody: "{\"message\":\"mongo: no documents in result\"}",
			username:     "not-existing-username",
		},
		{
			name:         "OK Happy Birthday",
			expectedCode: http.StatusOK,
			expectedBody: "{\"message\":\"Hello, username1! Happy birthday!\"}",
			username:     "username1",
		},
		{
			name:         "OK Tomorrow is Your Birthday",
			expectedCode: http.StatusOK,
			expectedBody: "{\"message\":\"Hello, username2! Your birthday is in 1 day(s)\"}",
			username:     "username2",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			setupMongoDBConnection()
			e := setupServer()
			w := performRequest(e, "GET", fmt.Sprintf("/hello/%s", tc.username), nil)
			assert.Equal(t, tc.expectedCode, w.Code)
			assert.Contains(t, tc.expectedBody, w.Body.String())
		})
	}
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
