package main

import (
	"fmt"
	"strconv"
)

func BoolToString() {
	isNew := true
	s := strconv.FormatBool(isNew)
	message := fmt.Sprintf("Is this a new user? %s", s)
	fmt.Println(message)
}

func InstringToInt() {
	s := "123"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Printf("string: %s\n int: %d\n", s, i)
	number := 100
	// Convert the integer to a string
	numberStr := strconv.Itoa(number)
	fmt.Println(numberStr)
	numberFloat := 23445221.1223356
	// Format the float with 3 decimal places
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f', 3, 64)
	fmt.Println(numberFloatStr)
	// Format the float with the smallest number of digits necessary to represent it
	numberFloatStr = strconv.FormatFloat(numberFloat, 'f', -1, 64)
	fmt.Println(numberFloatStr)
}
