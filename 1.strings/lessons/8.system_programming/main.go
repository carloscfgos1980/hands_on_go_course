package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	commands := map[string]func([]string){
		"CatchingSignals": func(_ []string) { CatchingSignals() },
		"ChildProcess":    func(_ []string) { ChildProcess() },
		"PassArguments":   RunPassArguments,
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

	run(os.Args[2:])
}

func RunPassArguments(args []string) {
	if len(args) < 1 {
		fmt.Println("usage: go run . PassArguments <argument>")
		return
	}

	PassArguments(args)
}

func commandNames(commands map[string]func([]string)) []string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
