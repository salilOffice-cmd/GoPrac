// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//In Go, the receiver of a function is the parameter that comes before the function name
//in a method definition. It specifies which type the method belongs to.
//Then you can call that method on the specified type.


// Struct
type Person struct {
	Name string
	Age  int
}

func (p Person) getName() string{
	return p.Name;
}

func (p Person) getAgeAfterTenYears() int{
	p.Age = p.Age + 10;
	return p.Age;
}

// This method will actually change the value of the struct being passed
func (p *Person) changeOriginalAge(){
	p.Age = 10;
}


// Map

// Define a type alias for a map
type MyMap map[string]int;

// Method for displaying the contents of MyMap
func (m MyMap) Display() {
    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}

func main() {
	

	// Calling struct methods
	var person5 = Person{"Bhoot", 100}
	fmt.Println("Person5 name:", person5.getName())	
	fmt.Println("Person5 age:", person5.getAgeAfterTenYears())


   
    	// Calling map methods
	myMap := MyMap{"apple": 5, "banana": 3, "orange": 7}
    	myMap.Display()

	
	// Calling methods that have a receiver pointer (mostly used)

	fmt.Println("Person5 before age:", person5.Age)
	person5.changeOriginalAge();
	fmt.Println("Person5 after age:", person5.Age)

}
