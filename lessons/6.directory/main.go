package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	commands := map[string]func([]string){
		"CheckLogFile":    func(_ []string) { CheckLogFile() },
		"CheckLogNotFile": func(_ []string) { CheckLogNotFile() },
		"ReadFile":        func(_ []string) { ReadFile() },
		"WriteFile":       func(_ []string) { WriteFile() },
		"WriteTempFile":   func(_ []string) { WriteTemporaryFile() },
		"FileCountLines":  func(_ []string) { CountLinesInFile() },
		"ReadLine":        RunReadLine,
		"CompareFiles":    func(_ []string) { fmt.Println(CompareFiles()) },
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

func RunReadLine(args []string) {
	if len(args) < 1 {
		fmt.Println("usage: go run . ReadLine <line-number>")
		return
	}

	lineNumber, err := strconv.Atoi(args[0])
	if err != nil || lineNumber < 0 {
		fmt.Println("line-number must be a non-negative integer")
		return
	}

	fmt.Println(ReadLine(lineNumber))
}

func commandNames(commands map[string]func([]string)) []string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
