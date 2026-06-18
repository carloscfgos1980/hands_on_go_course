package main

import (
	"fmt"
	"os"
	"strings"

	lessons "hands_on_go/strings/lessons/1.strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . [trim|substring]")
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "trim":
		lessons.StringsTrimSpace()
	case "substring":
		lessons.Substring()
	default:
		fmt.Println("unknown command: use trim or substring")
	}
}
