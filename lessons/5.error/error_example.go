package main

import "fmt"

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
