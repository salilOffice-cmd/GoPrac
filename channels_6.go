package main

import (
	"fmt"
	"time"
)

// In this example, we are dealing with three goroutines

func firstGoroutine(out chan<- string) {
	time.Sleep(2 * time.Second) // Simulate some work
	out <- "Data from first goroutine"
}

func secondGoroutine(in <-chan string, out chan<- string) {
	data := <-in // Wait for data from the previous goroutine
	fmt.Println("Received from first goroutine:", data)

	time.Sleep(2 * time.Second) // Simulate some work
	out <- "Data from second goroutine"
}

func thirdGoroutine(in <-chan string, done chan<- string) {
	data := <-in // Wait for data from the previous goroutine
	fmt.Println("Received from second goroutine:", data)

	// Perform some processing with the received data
	done <- "done"
}

func main() {
	first := make(chan string)
	second := make(chan string)
	done := make(chan string) // pass this channel to the last goroutine of the program

	go firstGoroutine(first)
	go secondGoroutine(first, second)
	go thirdGoroutine(second, done)

	//  Not a good practice
	//time.Sleep(5 * time.Second) // Wait for goroutines to finish

	// Good practice(creating a 'done' channel')
	var receviedFromDoneChannel string = <-done   // or receviedFromDoneChannel := <- done
	fmt.Println("Last message:", receviedFromDoneChannel)
}
