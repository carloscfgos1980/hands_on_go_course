package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type MyError struct {
	ShortMessage    string
	DetailedMessage string
	// Name string
	// Age int
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}
func CustomError() {
	err := doSomething()
	fmt.Print(err)
}
func doSomething() error {
	// Doing something here...
	return &MyError{ShortMessage: "Wohoo something happened!",
		DetailedMessage: "File cannot found!"}
}

func HandleErrorDate() {
	// Parsing a date with the wrong format will return an error
	parsedDate, err := time.Parse("2006", "2018")
	if err != nil {
		fmt.Println("An error occured", err.Error())
	} else {
		fmt.Println(parsedDate)
	}
	parsedDate, err = time.Parse("2006", "2018 abc")
	if err != nil {
		fmt.Println("An error occured", err.Error())
	} else {
		fmt.Println(parsedDate)
	}
}

func GoErrorPackage() {
	_, err := doSomething2()
	if err != nil {
		fmt.Println(err)
	}
}
func doSomething2() (string, error) {
	return "", errors.New("Something happened.")
}

func ErrorLogging() {
	log_file, err := os.Create("log_file")
	if err != nil {
		fmt.Println("An error occured...")
	}
	defer log_file.Close()
	log.SetOutput(log_file)
	log.Println("Doing some logging here...")
	log.Fatalln("Fatal: Application crashed!")
}

func ErrorPanic() {
	panic("Something went wrong!")
}

func DeferPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	writeSomething()
}

func writeSomething() {
	panic("Write operation error")
}

func sayHello() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	writeSomething()
}

func RecoverPanic() {
	sayHello()
	fmt.Println("After the panic was recovered!")
}
