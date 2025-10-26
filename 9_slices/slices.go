package main

import (
	"fmt"
	"slices"
)

// slice - dynamic sized array
// most used construct in go
// + usefull methods - append, copy, len, cap
func main() {
	// how to declare slice
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums)        // opt - [] nill
	fmt.Println(nums == nil) // opt - true

	// get length
	fmt.Println(len(nums)) // opt - 0

	// The make built-in function allocates and initializes an object of type slice, map, or chan (only)
	var numsUsingMake = make([]int, 2, 5) // initial 2 size we can add more later
	// above third parameter is capacity if not given capacity will be same as length
	fmt.Println(numsUsingMake) // opt - [0 0]

	// capacity - maximum number of elements can fit but since it is slice it can have more elements
	fmt.Println(cap(nums))

	// to add elements to slice
	numsUsingMake = append(numsUsingMake, 4)
	fmt.Println(numsUsingMake) // opt - [0 0 4]

	// Third way to create slice
	nums3 := []int{} // default capacity will be 1
	fmt.Println(nums3)

	// add element at index
	// nums3[0] = 5
	nums3 = append(nums3, 5)
	fmt.Println(nums3) // opt - [5]

	//copy slice
	nums4 := make([]int, len(nums3))

	fmt.Println(nums4)

	copy(nums4, nums3)
	fmt.Println(nums4) // opt - [5]

	// slice operator

	var numsSlice = []int{1, 2, 3, 4}
	fmt.Println(numsSlice[0:2]) // prints slice including 0 excluding 2
	fmt.Println(numsSlice[1:])  // prints slice from index 1 to end
	fmt.Println(numsSlice[:3])  // prints slice from start to index 3 excluding 3
	fmt.Println(numsSlice[:])   // prints whole slice

	// we also have slices package - inbuilt standard library package
	var slicenums1 = []int{1, 2}
	var slicenums2 = []int{1, 2}

	fmt.Println(slices.Equal(slicenums1, slicenums2)) // compares slices returns true if equal

	// 2d slices
	var nums2d = [][]int{{1}, {1, 2}, {1, 2, 3}}
	fmt.Println(nums2d)
}
