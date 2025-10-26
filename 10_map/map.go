package main

import (
	"fmt"
	"maps"
)

// MAPS - hash, object ,dictionary stored in key value pairs
func main() {
	// create map
	m := make(map[string]int)
	// map[type of key]type of value

	//setting an element
	m["a"] = 1
	m["b"] = 2

	//get an element
	fmt.Println(m["a"])

	// what if we access a key that does not exist
	fmt.Println(m["c"]) // returns zero value of int which is 0
	// if key doesn't exist in map it will return zero value of type

	// length of map
	fmt.Println("length of map is:", len(m))

	// deleting an element
	delete(m, "b")      //delets element
	fmt.Println(m["b"]) // returns zero value of int which is 0

	// to clear a map
	clear(m)
	fmt.Println("length of map after clearing is:", len(m))

	// creating map without using make
	m1 := map[string]int{
		"x": 10,
	}
	fmt.Println("map m1:", m1)

	// check if key exists
	k, ok := m1["x"] // go also provides few more thing along with value
	fmt.Println("value of k is:", k)
	// ok is boolean which tells if key exists or not
	// first value is value of key second is boolean of key existenceff
	if ok {
		fmt.Println("all Ok")
	} else {
		fmt.Println("not ok")
	}

	map1 := map[string]int{"x": 10, "y": 20}
	map2 := map[string]int{"x": 10, "y": 20}

	// to check if two maps are equal
	// we have maps package that is similar like slices pacakge has all maps related methods
	if maps.Equal(map1, map2) {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}

}
