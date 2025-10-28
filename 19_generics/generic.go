package main

import "fmt"

//Generic = Generics in programming refer to the ability to define classes, interfaces, and methods with type parameters, allowing them to operate on a variety of data types without needing to specify the exact type in advance

// generics  in go introduced in 1.18 version

// func printSlice(items []int) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }

//what if we have above int slice
// what if we have above float slice
// what if we have above bool slice

// For all cases we than need to create another methods
// This can be resolved using generic

func printSlice[T any](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

// generics can be implementd as above we pass [T any] after function name
// and use it in parameter
// it can be T , A ,B anything but generally T is prefereed

// now above function will work with all type

// if we want specific few types we can also do this [T string | int | float] etc
// we can also use interface {} instead of any it works same way

func printIntStringSlice[T int | string](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

// we can  also use geneic in struct

// if we create string stack elements will be all of string type
// same if we create int stack it will have all elements of int type
type stack[T any] struct {
	elements []T
}

// we can also pass multiple generics

func printSlicewithName[T comparable, V string](items []T, name V) {
	for _, v := range items {
		fmt.Println(v)
	}
	fmt.Println(name)
}

func main() {
	printSlice([]int{1, 2, 3})

	printSlice([]string{"hello", "world"})

	printIntStringSlice([]int{1, 2, 3})

	printIntStringSlice([]string{"hello", "world"})

	// printIntStringSlice([]bool{false, true}) this will not work because we mentioned int and string type in generic

	myStringStack := stack[string]{
		elements: []string{"go lang"},
	}
	myIntStack := stack[int]{
		elements: []int{1, 2, 3, 4},
	}

	fmt.Println(myStringStack)
	fmt.Println(myIntStack)

	printSlicewithName[int, string]([]int{1, 2}, "generic")
}

// we can  also use comparable type which will include all comparable types

// func printComparableSlice[T comparable](items []T) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }
