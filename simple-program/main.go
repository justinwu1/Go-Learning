package main

import (
	"fmt"
	"log"
	"time"
)



// definie interface， just like C#, its use to loose couple.
type Animal interface{
	Talk() string 
	Legs() int

}
type Dog struct{
	Name string 
	Breed string
}

type Gorilla struct{
	Name  string
	Color string 
	Teeth int 
}
func main() {
	fmt.Println("Hello world")
	var name string = "Justin"
	fmt.Println(name)
	functionVar,functionVar2 := saySomething()
	fmt.Println(functionVar)
	fmt.Println(functionVar2)

	var color string = "green"
	changeThroughPointer(&color)
	fmt.Println(color)
	
	newUser := User{
		FirstName: "Justin",
		LastName: "Wu",
	}
	log.Println(newUser.FirstName)
	log.Println(newUser.printFirstName())

	// create a map, call the make function and specfiy the type of the key&value, even struct 
	myMap := make(map[string]string)
	myMap["dog"] = "Justin"

	// Create slices(array) 
	var mySlice []string
	mySlice = append(mySlice, "Justin")
	mySlice = append(mySlice, "John")

	// for loop syntax, range function = length, _ = blank identifier 
	for i:=0; i <= 10; i++{
		log.Println(i)
	}
}
// Return 2 value and assign to two variable
func saySomething()(string,string){
	return "something","else"
}

// Change value with its pointer
func changeThroughPointer(s *string){
	newValue := "red"
	*s = newValue
}

// Types and  struct
type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

// Assign function to struct 
func(receiver *User)printFirstName()string{
	return receiver.FirstName
}
