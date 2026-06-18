package lessons

import (
	"fmt"
	"strings"
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
