package main

import "fmt"

func main() {
	// STRING VARIABLE
	var name string = "Radhey" //in go lang we need to use declared variables
	fmt.Println(name)
	// Above can also be written as
	var name2 = "Radhey Mugdal" // if wfe don't specify type go lang will infer its style autometically
	// this can  be used with any data type like integer, boolean, float etc
	fmt.Println(name2)

	// INT VARIABLE
	var age int = 21
	fmt.Println(age)

	// SHORT HAND DECLARATION
	language := "Golang" // only works inside functions
	// in same time it will declares and assigns value
	fmt.Println(language)

	//When we need to assign type to variable
	var name3 string // if we assign type string it will then take only string values
	name3 = "Radhey Mugdal"
	fmt.Println(name3)

	// FLOAT VARIABLE
	var height float32 = 5.9
	fmt.Println(height)

}
