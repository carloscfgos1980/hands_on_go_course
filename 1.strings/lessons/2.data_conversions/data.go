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

func StringToInt() {
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

func StringToBool() {
	s := "true"
	b, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
		return
	}
	fmt.Printf("string: %s\n bool: %t\n", s, b)
	isNew := "j"
	isNewBool, err := strconv.ParseBool(isNew)
	if err != nil {
		fmt.Println("failed")
	} else {
		if isNewBool {
			fmt.Print("IsNew")
		} else {
			fmt.Println("Not new")
		}
	}
}

func StringToFloat() {
	s := "3.14"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}
	fmt.Printf("string: %s\n float: %f\n", s, f)
	f, err = strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}
	fmt.Printf("string: %s\n float: %f\n", s, f)
}

func ByteToString() {
	helloWorldByte := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111,
		114, 108, 100}
	fmt.Println(string(helloWorldByte))
}

func StringToByte() {
	helloWorld := "Hello, World"
	helloWorldByte := []byte(helloWorld)
	fmt.Println(helloWorldByte)
}
