package main

import "fmt"

// numbered sequence of specific length

func main() {
	// default values are 0 for int, 0 for float, "" for string, false for bool
	var arr [4]int

	arr[0] = 1
	fmt.Println(arr[0])

	// how to get length
	fmt.Println(len(arr))

	var names [3]string
	names[0] = "Radhey"
	fmt.Println(names)

	// assign value while declaration
	nums := [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	//2d arrays

	arr2d := [2][3]int{{1, 2, 4}, {5, 6, 7}}
	fmt.Println(arr2d)

	// array is fixed size
	// if we know size at compile time we can use arrays  gives constant time access
	// if we want dynamic size we can use slices
}
