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

func ReadFile() {
	contentBytes, err := os.ReadFile("names.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var contentStr string = string(contentBytes)
	fmt.Println(contentStr)
}

func WriteFile() {
	content := "Hello, World!\nThis is a test file.\n"
	err := os.WriteFile("hello_world.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")
}
