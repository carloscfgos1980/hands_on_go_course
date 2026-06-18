package lessons

import (
	"fmt"
	"strings"
)

func StringsTrimSpace() {
	greetings := "\t Hello, World "
	fmt.Printf("lent: %d\n greeting: %s", len(greetings), greetings)
	trimmed := strings.TrimSpace(greetings)
	fmt.Printf("lent: %d\n greeting: %s", len(trimmed), trimmed)
}

func Substring() {
	greeting := "Hello, World and Mars"
	helloWorld := greeting[0:12]
	fmt.Println(helloWorld)
	mars := greeting[16:]
	fmt.Println(mars)
}
