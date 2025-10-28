package main

import "fmt"

// enum are enumerated types (custom type)

// func changeOrderStatus(status string) {
// 	fmt.Println("changing order status to ", status)
// }

// above status are string so it can  be anything string
// but suppose if we want only few string like "confirmed" , "placed", "delivered"
// above three string is custom type

// Go does't have any specific enum like other language
// We use const and type to create const as below

type OrderStatus string

// const (
// 	Received  OrderStatus = iota // iota is integer which will get increment for first time it will give 0
// 	Confirmed                    // this will be 1
// 	Prepared                     // this will be 2
// 	Delivered                    // this will be 2
// )

// if we want to use string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	// changeOrderStatus(Confirmed) // status to 1
	changeOrderStatus(Received)
}
