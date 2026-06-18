package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "\t Hello, World "
	fmt.Printf("lent: %d\n greeting: %s", len(greetings), greetings)
	trimmed := strings.TrimSpace(greetings)
	fmt.Printf("lent: %d\n greeting: %s", len(trimmed), trimmed)
}
