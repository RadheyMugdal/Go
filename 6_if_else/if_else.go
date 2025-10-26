package main

import "fmt"

func main() {
	age := 18
	if age < 18 {
		fmt.Println("You are a minor")
	} else if age == 18 {
		fmt.Println("You are just became an adult")
	} else {
		fmt.Println("You are an adult")
	}

	var role = "admin"
	var hasPermission bool = false

	if role == "admin" || hasPermission {
		fmt.Println("You can access admin panel")
	}

	if role == "admin" && hasPermission {
		fmt.Println("You can access admin panel")
	}

	// declare variable directly insider if

	if age := 15; age < 18 {
		fmt.Println("You are a minor", age)
	} else if age == 18 {
		fmt.Println("You are just became an adult", age)
	} else {
		fmt.Println("You are an adult", age)
	}

	// go does not have ternary, you will have to use nnormal if else
	// (X > Y) ? X : Y  ---- not valid in go
}
