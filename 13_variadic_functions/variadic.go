package main

import "fmt"

// variadic function - variadic function are function where you can pass any number of parameters like shown below in println
// fmt.Println(1, 2, 3, "hello")

// How to create it

func sum(nums ...int) int { // this will receive n number of integers as parametrs
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

//to receive any type of parameters
func printAll(val ...interface{}) {
	// to spread all values we do val...
	fmt.Println(val...)
}

func main() {
	fmt.Println(1, 2, 3, "hello")

	fmt.Println(sum(1, 2, 23, 4, 5, 3)) // opt - 38
	printAll("hello", 1, "world")
}
