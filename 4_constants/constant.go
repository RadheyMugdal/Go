package main

import "fmt"

// constants can also be declared outside of function
const age = 21

//variables can also be declared outside of function
var country = "India"

//but short hand declaration can not be used outside of function
// language:="Golang" // this will give error

func main() {
	// Constants are values that can not be changed
	// we must need to assign value at the time of declaration
	const name = "Radhey Mugdal"

	//name = "New Name" // This will give error as constants can not be changed

	fmt.Println(name)

	// If we want to have multiple constants in one go we can do like below
	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(port, host)
}
