package main

import (
	"fmt"
)

// below is how functions can be declared
// we can also declare below parameter like this (a,b int) this will make all parameters of type int
func add(a int, b int) int {
	return a + b
}

// In Go function can return multiple values
// below how we can return multiple values in return type we mention all return values types
func getLanguages() (string, string, int) {
	return "Go", "Javascript", 23
}

// In Go functions are first class citizen function we can also assign funtions to variables, we can also pass function as argument in another function

func processIt(fn func(a int) int, b int) int {
	result := fn(3)
	return result + b
}

// main is also funcion this is entry point of whole go program
func main() {

	result := add(2, 3)
	fmt.Println(result)
	lang1, lang2, number := getLanguages()
	fmt.Println(lang1, lang2, number)
	fmt.Println(getLanguages())

	res := processIt(func(a int) int {
		return a - 1
	}, 5)

	fmt.Println(res)

	// anonymous function without name
	// and assign function to variable
	fn := func() string {
		return "new"
	}
	fn()

}
