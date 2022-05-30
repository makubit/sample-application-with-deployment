package main

import (
	"fmt"
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
