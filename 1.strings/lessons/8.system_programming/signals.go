package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func CatchingSignals() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		fmt.Println(sig)
		fmt.Println("Signal captured and processed...")
		done <- true
	}()
	fmt.Println("Waiting for signal")
	<-done
	fmt.Println("Exiting the application...")
}

func ChildProcess() {
	lsCommand := exec.Command("ls")
	output, _ := lsCommand.Output()
	lsCommand.Run()
	fmt.Println(lsCommand.Process.Pid)
	fmt.Println(string(output))
}
