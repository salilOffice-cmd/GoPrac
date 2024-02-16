// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}



func main() {

	// STRUCT
	// Initializing a struct using a struct literal
	var person1 Person = Person{"Alex", 30}
	person2 := Person{"Salil", 22}
	fmt.Println(person1)
	fmt.Println(person2.Name, person2.Age)

	// Initializing a struct using named fields
	person3 := Person{
		Name: "Ram",
		Age:  25,
	}
	fmt.Println(person3)

	// Declaring a struct without initializing
	var person4 Person
	person4.Name = "Rohit"
	person4.Age = 23
	fmt.Println(person4)
	

	// MAP

	// 1. Creating a map
	// Initiliazing a map
	var map1 = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(map1)
	fmt.Println(map1[1]) // accessing value of a key

	// using make() function
	var map2 = make(map[string]string) // The map is empty now
	map2["salil"] = "Nagpur"
	map2["aanchal"] = "koradi"
	fmt.Println(map2)

	// 2. Deleting a key-value pair
	delete(map2, "salil")
	fmt.Println(map2)

	// 3. Iterating over a map
	//for k, v := range map1 {
	//	fmt.Printf("%v : %v, ", k, v)
	//}

	// 4. Checking existence of a key value pair
	val1, ok1 := map1[1]  // Checking for existing key and its value
	val2, ok2 := map1[11] // Checking for non-existing key and its value
	val3, ok3 := map1[1]  // Checking for existing key and its value
	_, ok4 := map1[123]   // Only checking for existing key and not its value(imp)

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}
