package main

import (
	"fmt"
	"time"
)

func main() {
	// switch statement is used to select one of many code blocks to be executed

	// simple switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
		// break is not required in go lang switch case
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Wrong day")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	//type switch

	whomi := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("default")
		}
	}
	whomi("hello")
}
