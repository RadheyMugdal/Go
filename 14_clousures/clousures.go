package main

import "fmt"

// clousures - closure is function that remembers its environment in which it was created
// even after that environment has finished executing
// this means function can still access variables from its outer function
//even after function has returned

func counter() func() int {
	counter := 0

	return func() int {
		counter += 1
		return counter
	}
}
func main() {
	increment := counter()

	fmt.Println(increment()) // opt- 1
	fmt.Println(increment()) // opt -2 it will remember last state of opt since we took it from outer function
	// generally if function is removed from call stack it loose that counter state if it is in same function but we are
	// creating it in outer function
}
