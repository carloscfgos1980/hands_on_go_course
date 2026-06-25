package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func StringsTrimSpace() {
	greetings := "\t Hello, World "
	fmt.Printf("lent: %d\n greeting: %s\n", len(greetings), greetings)
	trimmed := strings.TrimSpace(greetings)
	fmt.Printf("lent: %d\n greeting: %s\n", len(trimmed), trimmed)
}

func Substring() {
	greeting := "Hello, World and Mars"
	helloWorld := greeting[0:12]
	fmt.Println(helloWorld)
	mars := greeting[16:]
	fmt.Println(mars)
}

func ReplaceString() {
	greeting := "Hello, World and Mars"
	newGreeting := strings.Replace(greeting, "Mars", "Venus", 1)
	fmt.Println(newGreeting)
}

func ReplaceParts() {
	greeting := "Hello, World and Mars"
	// Replace 2 spaces with hyphens
	newGreeting := strings.Replace(greeting, " ", "-", 2)
	fmt.Println(newGreeting)
	// Replace all spaces with hyphens
	newGreeting = strings.ReplaceAll(greeting, " ", "-")
	fmt.Println(newGreeting)
}

func EscapingCharacters() {
	helloWorld := "Hello World, this is Carlos \"The Great\""
	fmt.Println(helloWorld)
	helloWorld = "Hello World, this is Carlos \"The Great\" \t\n Hello Again"
	fmt.Println(helloWorld)
	helloWorld = "Hello World, this is Carlos \"The Great\" \\tHello Again"
	fmt.Println(helloWorld)
}

func Capitalizing() {
	greeting := "hello, world"
	// Capitalize the first letter of the greeting
	capitalized := strings.ToUpper(greeting[:1]) + greeting[1:]
	fmt.Println(capitalized)
	// Capitalize the first letter of each word in the greeting
	title := cases.Title(language.English).String(greeting)
	fmt.Println(title)
	// Capitalize the entire greeting
	upper := strings.ToUpper(greeting)
	fmt.Println(upper)
}
