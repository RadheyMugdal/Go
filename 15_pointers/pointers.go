package main

import "fmt"

// pointer - pointer is address of memory location of variable

func changeNum(num int) {
	num = 5
	fmt.Println("In ChangeNum", num)
}

// below takes pointer of that value means we pass memory address and change actual variable
func changeActualNum(num *int) {
	*num = 5 // put it in pointer's location we de referenced it
	fmt.Println(*num)
}

func main() {
	num := 1
	changeNum(num)                               //num pass by value
	fmt.Println("After change num in main", num) // this will print 1 only not 5
	// because we passed num by value
	// pass by value means copy of our actual variable

	// We can pass variable by reference means actual location of our variable not copy
	//actual variable
	// this can be done by pointer since it is having actual memory address of variable

	fmt.Println("Memory address of num", &num) // this will print memory address

	changeActualNum(&num)
	fmt.Println("After change num by using pointer change num", num) // opt this will print 5 not 1 it will change actual variable
	//pass by reference
}
