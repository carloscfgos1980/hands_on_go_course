package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	lessons "hands_on_go/strings/lessons/1.strings"
)

func main() {
	commands := map[string]func(){
		"Replace":   lessons.ReplaceString,
		"TrimSpace": lessons.StringsTrimSpace,
		"Substring": lessons.Substring,
	}

	if len(os.Args) < 2 {
		fmt.Printf("usage: go run . [%s]\n", strings.Join(commandNames(commands), "|"))
		return
	}

	run, ok := commands[os.Args[1]]
	if !ok {
		fmt.Printf("unknown command %q. available: %s\n", os.Args[1], strings.Join(commandNames(commands), ", "))
		return
	}

	run()
}

func commandNames(commands map[string]func()) []string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
