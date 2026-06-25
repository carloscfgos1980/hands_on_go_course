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

func AddTime() {
	// Get the current time
	current := time.Now()
	// Add 5 hours to the current time
	nextFiveHours := current.Add(5 * time.Hour)
	fmt.Println("Next Five Hours:", nextFiveHours.Format("2006-01-02 15:04:05"))
	// Add 30 minutes to the current time
	nextHalfHour := current.Add(30 * time.Minute)
	fmt.Println("Next Half Hour:", nextHalfHour.Format("2006-01-02 15:04:05"))
	// Add 45 seconds to the current time
	next45Seconds := current.Add(45 * time.Second)
	fmt.Println("Next 45 Seconds:", next45Seconds.Format("2006-01-02 15:04:05"))
}

func SubtractTime() {
	// Get the current time
	current := time.Now()
	// Subtract 5 hours from the current time
	lastFiveHours := current.Add(-5 * time.Hour)
	fmt.Println("Last Five Hours:", lastFiveHours.Format("2006-01-02 15:04:05"))
	// Subtract 30 minutes from the current time
	lastHalfHour := current.Add(-30 * time.Minute)
	fmt.Println("Last Half Hour:", lastHalfHour.Format("2006-01-02 15:04:05"))
	// Subtract 45 seconds from the current time
	last45Seconds := current.Add(-45 * time.Second)
	fmt.Println("Last 45 Seconds:", last45Seconds.Format("2006-01-02 15:04:05"))
}

func TimeDifference() {
	// Get the current time
	current := time.Now()
	// Create a future time 2 hours from now
	future := current.Add(2 * time.Hour)
	// Calculate the difference between the future time and the current time
	diff := future.Sub(current)
	fmt.Println("Time Difference:", diff)
}

func ParseDate() {
	str := "2018-08-08T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"
	noDateStrig := "lalalala"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Println("Parsed Date:", t)
	// Attempt to parse an invalid date string
	_, err = time.Parse(layout, noDateStrig)
	if err != nil {
		fmt.Println("Error parsing invalid date string:", err)
	}
}
