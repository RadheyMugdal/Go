package main

// time package provided by go
import (
	"fmt"
	"time"
)

// Struct is custom data structure which like we create classes in other languages and creates object and use it
// But in go we don't have class , We can use struct

// Let create order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // neno second precision
}

// we can also assign some methods to struct just like we assign to class it other languages
// by putting (o order) we will attach this method to that struct
// just like we do in other language for example getter setter
// this method wiil have access to order struct variables
// if we use (name type) (o order ) than this will be pass by value
// if we do this than it will channge copy of  this order not actual one
// to change actual one we need to use pointer and use pass by reference
// as we learnt in previews tut
func (o *order) changeStatus(status string) {
	o.status = status
	// why we haven't use dereference thing eg *o.status ?
	// because struct will autometically does that we dont need to
}

func (o order) getAmount() float32 {
	return o.amount
}

//  4
//Constructers In Go
// A constructor is a special method in object-oriented programming that is automatically called when an object is created to initialize it

// we don't have anything like constructor
// but we do it as below
// we create function and create instance of that struct using that
// this way we can have something like constructor and we can also put
// any initial code that we usually write in constructor here in that function

type user struct {
	name string
}

func newuser(name string) *user {
	// initial setup goes here
	myOrder := user{
		name: name,
	}
	return &myOrder
}

// 5
// STRUCT EMBEDDING
// we can also use another struct in struct for referece
// for example in order struct we want to referece customer struct we can do it as below

type customer struct {
	name  string
	phone string
}

type orderWithCustomer struct {
	id       string
	status   string
	amount   float32
	customer // this is referenced struct from customer this is struct embedding
}

func main() {
	// now we will create instance of that struct just like we create for class in other language
	myOrder := order{
		id:     "1",
		amount: 50,
		status: "Placed",
	}
	// if we want to add or update later on that can also be done
	myOrder.createdAt = time.Now()
	myOrder.status = "delivered"

	// print myOrder
	fmt.Println(myOrder)

	myOrder2 := order{
		id:        "2",
		amount:    54,
		status:    "delivered",
		createdAt: time.Now(),
	}
	fmt.Println(myOrder2)
	// we can use change status here
	myOrder2.changeStatus("Cancelled")
	fmt.Println(myOrder2)

	fmt.Println(myOrder2.getAmount())

	//what if we don't set feild while creating instance

	myOrder3 := order{
		id:     "1",
		status: "delivered",
	}
	fmt.Println(myOrder3)

	// If we dont set feild while creating instance it will by default assign its zero value of perticular type
	// int --> 0
	// string -->""
	// bool --> false

	user := *newuser("radhey")
	fmt.Println(user)

	// 4
	// Since we also need to use struct only for one time
	// For example we create objects
	// We can also do it in go
	language := struct {
		name   string
		isGood bool
	}{
		"GoLang", true,
	}

	fmt.Println(language)

	// 5
	order4 := orderWithCustomer{
		id:     "1",
		status: "processing",
		amount: 12.99,
		customer: customer{ // if we dont initialze this customer it will gona have empty value struct like for here { }
			name:  "radhey",
			phone: "3234324324",
		},
	}

	fmt.Println(order4)
}
