// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second) // Simulate some work
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(1 * time.Second) // Simulate some work
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printNumbers() // Start a new goroutine to execute printNumbers concurrently
	go printLetters() // Start another goroutine to execute printLetters concurrently

	time.Sleep(10 * time.Second) // Wait for goroutines to finish (not a good practice, just for this example)
	fmt.Println("Main function exits")
}
