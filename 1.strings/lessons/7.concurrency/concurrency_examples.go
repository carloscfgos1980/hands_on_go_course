package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintNames() {
	names := []string{"tarik", "john", "michael", "jessica"}
	for _, name := range names {
		time.Sleep(1 * time.Second)
		fmt.Println(name)
	}
	ages := []int{1, 2, 3, 4, 5}
	for _, age := range ages {
		time.Sleep(1 * time.Second)
		fmt.Println(age)
	}
}

func PrintNamesConcurrently() {
	go func() {
		names := []string{"tarik", "john", "michael", "jessica"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
	}()
	go func() {
		ages := []int{1, 2, 3, 4, 5}
		for _, age := range ages {
			time.Sleep(1 * time.Second)
			fmt.Println(age)
		}
	}()
	time.Sleep(10 * time.Second)
}

func PrintNameConcurrencyChannel() {
	nameChannel := make(chan string)
	go func() {
		names := []string{"tarik", "john", "michael", "jessica"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
		nameChannel <- ""
	}()
	<-nameChannel
}

func PassingDataChannel() {
	nameChannel := make(chan string, 5)
	done := make(chan string)
	go func() {
		names := []string{"tarik", "michael", "gopi", "jessica"}
		for _, name := range names {
			// doing some operation
			fmt.Println("Processing the first stage of: " + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()
	go func() {
		for name := range nameChannel {
			fmt.Println("Processing the second stage of: " + name)
		}
		done <- ""
	}()
	<-done
}

func ConcurrentFinish() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Hello World")

		}()
	}
	wg.Wait()
}

func ResultConcurrentFunctions() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hello from channel1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Hello from channel2"
	}()
	var result string
	for {
		select {
		case result = <-channel1:
			fmt.Println(result)
		case result = <-channel2:
			fmt.Println(result)
		default:
			return
		}
	}
}
