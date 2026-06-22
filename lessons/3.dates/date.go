package main

import (
	"fmt"
	"time"
)

func TimeNow() {
	// Get the current time
	current := time.Now()
	// Print the current time in the default format
	fmt.Println(current.String())
	// Print the current time in a custom format
	fmt.Println("MM-DD-YYYY :", current.Format("01-02-2006"))
	fmt.Println("YYYY/MM/DD :", current.Format("2006/01/02"))
	fmt.Println("Day Month Year :", current.Format("02 Jan 2006"))
	// Print the current time in a custom format with time
	fmt.Println("YYYY-MM-DD hh:mm:ss", current.Format("2006-01-02 15:04:05"))
}

func AddDate() {
	// Get the current time
	current := time.Now()
	// Add 1 day to the current time
	tomorrow := current.AddDate(0, 0, 1)
	fmt.Println("Tomorrow:", tomorrow.Format("2006-01-02"))
	// Add 1 month to the current time
	nextMonth := current.AddDate(0, 1, 0)
	fmt.Println("Next Month:", nextMonth.Format("2006-01-02"))
	// Add 1 year to the current time
	nextYear := current.AddDate(1, 0, 0)
	fmt.Println("Next Year:", nextYear.Format("2006-01-02"))
}

func SubtractDate() {
	// Get the current time
	current := time.Now()
	// Subtract 1 day from the current time
	yesterday := current.AddDate(0, 0, -1)
	fmt.Println("Yesterday:", yesterday.Format("2006-01-02"))
	// Subtract 1 month from the current time
	lastMonth := current.AddDate(0, -1, 0)
	fmt.Println("Last Month:", lastMonth.Format("2006-01-02"))
	// Subtract 1 year from the current time
	lastYear := current.AddDate(-1, 0, 0)
	fmt.Println("Last Year:", lastYear.Format("2006-01-02"))
}
