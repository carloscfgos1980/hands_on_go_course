package main

import (
	"fmt"
	"os"
)

func CheckLogFile() {
	if _, err := os.Stat("log.txt"); err == nil {
		fmt.Println("Log.txt file exists")
	}
}

func CheckLogNotFile() {
	// Check if log.txt file exists
	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		fmt.Println("Log.txt file does not exist")
	} else {
		fmt.Println("Log.txt file exists")
	}

	// Check if log1.txt file exists
	if _, err := os.Stat("log1.txt"); os.IsNotExist(err) {
		fmt.Println("Log1.txt file does not exist")
	} else {
		fmt.Println("Log1.txt file exists")
	}
}
