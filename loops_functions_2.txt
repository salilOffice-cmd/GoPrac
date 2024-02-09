// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, 世界")
	}

	// Iterating over an array
	var fruits = [...]string{"apple", "orange", "banana"}
	fmt.Println(fruits)

	for i, val := range fruits {
		fmt.Printf("%v - %v\n", i, val)
	}

	// Creating and calling a function
	var addition int = Add(22, 2)
	fmt.Println(addition)
	
}
