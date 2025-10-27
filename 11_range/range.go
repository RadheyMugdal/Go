package main

import "fmt"

// Range in Go is used to iterate over elements in various data structures like arrays, slices, maps, strings, and channels.
func main() {
	nums := []int{1, 2, 3, 5, 6}

	// normal looping through for looop
	for i := 0; i < len(nums); i++ {
		fmt.Println("index:", i, "value:", nums[i])
	}

	// looping through range
	// range returns two values one is index and other is value at that index
	for k, i := range nums {
		fmt.Println(k, ":", i)
	}

	// iterate over map using range
	m := map[string]string{
		"fname": "john",
		"lname": "doe",
	}

	// if we take only single value like val:=range m it will give key
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	// range can also be used on string
	// first parameter returned is starting byte of runef
	for _, c := range "Go lang" {
		// it will print uni code of each character
		fmt.Println(c)

		// we can convert to string using string(c)
	}

	
}
