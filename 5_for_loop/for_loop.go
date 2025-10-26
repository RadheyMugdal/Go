package main

import "fmt"

// go lang has only one loop which is for loop
// go also doesn't have while loop or do while loop etc
// but we can use for loop to achieve same functionality

func main() {
	//  how to implement while loop using for loop
	i := 1

	for i <= 5 {
		fmt.Println("Value of i is:", i)
		i++
	}

	// infinite loop

	// for {
	// 	fmt.Println(1)
	// }

	//classic for loop
	for j := 0; j < 3; j++ {
		fmt.Println("Value of j is:", j)
	}

	// we also have continue and break statements
	for j := 0; j < 3; j++ {
		if j == 1 {
			continue
		}
		fmt.Println("Value of j is:", j)
	}
	for j := 0; j < 3; j++ {
		if j == 1 {
			break
		}
		if j == 2 {
			break
		}
		j++
		fmt.Println("Value of j is:", j)
	}

	// range

	for k := range 3 { // 0 1 2
		fmt.Println("Value of k is:", k)
	}
}
