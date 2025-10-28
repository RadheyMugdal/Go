package main

import "fmt"

// type payment struct{}

// func (p payment) makePayment(amount float32) {
// 	razorpayPaymentGW := razorpay{}
// 	razorpayPaymentGW.pay(amount)
// }

// type razorpay struct {
// }

// func (r razorpay) pay(amount float32) {
// 	// logic to make payment
// 	fmt.Println("making payment using razorpay", amount)

// }

// func main() {
// 	payment := payment{}
// 	payment.makePayment(21)
// }

// now if we want to have stripe instead
// than we need to change methods that we created
// we need to create new struct and change previous methods
// this is not good practice
// violates open close priciple - method should be open for upgrade but close for modification

// We This bad practice can be avoided using interface

// type payment struct {
// 	gateway stripe
// }
// type stripe struct{}

// func (p payment) makePayment(amount float32) {
// 	p.gateway.pay(amount)
// }

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

// func main() {
// 	stripeInstance := stripe{}
// 	newPayment := payment{gateway: stripeInstance}
// 	newPayment.makePayment(12)
// }

// we have added gateway feild in payment
// Above also doesn't solves our major problem
// if we want to have razorpay instead now
// than we need to change gateway type also than use repective pay method or  new payment gateway
// basically we need to change struct

// For example in testing we want to pass some fake gateway we cant do that here

// This can be resolved using interface
// Interface =An interface is a type that defines a set of method signatures,
// but it doesn't provide any implementation for these methods
// it serves as a contract , outlining a specific behaviour that other types can  fulfill

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

type stripe struct {
}
type razorpay struct {
}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}
func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

func main() {
	razorpay := razorpay{}

	payment := payment{
		gateway: razorpay, // we can pass stripe also since it implemennts that method described in interface
		// and go will implecitely detect that and allow us to use razorpay and stripe which implements methods declare in interface
		// Go will identify if any type that we are using implements or follows interface 
		// By seeing if it implements methods defined in interface 
		// if it will not than it throws error
	}
	payment.gateway.pay(121)
}
