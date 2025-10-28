package main

import (
	"fmt"
	"time"
)

//Channels -  Channels are the pipes that connect concurrent goroutines
// channels used for communication between different go routines

func processNum(numChan chan int) {
	fmt.Println("Processing number", <-numChan)
}

func processInfiniteNum(numChan chan int) {

	// if we are using range than we dont need <- this thing
	for num := range numChan {
		fmt.Println("Processing number", num)
	}
}

func sum(result chan int, num1 int, num2 int) {
	ans := num1 + num2
	result <- ans
}
func main() {
	// how to create channnel
	// we create it using make using chan and type of chan like below
	// messageChan := make(chan string)

	// how to send data to channel
	// messageChan <- "Ping" // we are sending data to channel

	// how to receive data
	// msg := <-messageChan // here we are receiving data from channel <- comes before channel
	// fmt.Println(msg)

	//above code will give error
	// ERROR
	// all go routines are asleep -deadlock
	// deadlock is operating system concept
	// whenever multiple processes running concurently it will hold resources while processing
	// and othher process will wait till first one release resources
	// and let suppose if first process is waiting for some resource it will wait if second process is holding that resource than both will than wait for each other
	// both are stuck in infinite waiting this is called deadlock

	// Why above code giving error?
	// 	Because:

	// You had no goroutine running concurrently.

	// The main goroutine tried to send into the channel messageChan <- "Ping".

	// But no one was receiving from that channel at that time — so the main goroutine blocked forever waiting for a receiver.

	// Since there’s only one goroutine (main), nothing else can receive → deadlock.
	// we are using
	// msg := <-messageChan  this is blocking , channels are blocking that's why it gives error

	// 	main creates a channel.

	// main tries to send (ch <- "Hi").

	// Channels are blocking — so it waits for someone to receive.

	// But there is no other goroutine running — the main goroutine itself can’t both send and receive at the same time.

	// So it waits forever → deadlock.

	numChan := make(chan int)
	go processNum(numChan)
	numChan <- 5

	time.Sleep(time.Second * 2)

	// main is also go routin and you can see we passed data between both of them using channel

	// how to send multiple data
	// this channel are used as queue many type
	// this can be done as below

	// go processInfiniteNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//receive channel data
	// below is how we can receive data from another channel

	result := make(chan int)
	go sum(result, 4, 3)
	res := <-result // this is also blocking that why we doesn't need to sleep

	fmt.Println(res)

}
