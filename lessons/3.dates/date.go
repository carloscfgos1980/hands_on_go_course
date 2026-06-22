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
}
