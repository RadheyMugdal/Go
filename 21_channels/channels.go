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

// Topic 5
// channel can be used to sync go routine
func task(done chan bool) {
	defer func() { // defer will make it run at very last
		done <- true
	}()
	fmt.Println("Processing task...")
}

//Topic 6

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("sending email to", email)
	}
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

	//Topic 5
	// channel can  also be used to syncronise go routing
	done := make(chan bool)
	go task(done)
	<-done // since this is blocking it will wait before completing go routing
	// will complete only after getting data

	// Topic 6
	// channel that we created using make are blocking
	// while receiving and sending data it will block the thread
	// but if we want to have something like queue which
	// is non blocking mostly. that can be created using buffered channel
	// UnBuffered channel = this are channel where we can sennd only one data at time and another can only be sent after someone receives that data
	// Above are not as we use in queue queue are not unbuffered
	// In Buffered channel we can send limited amount of data without blocking

	emailChan := make(chan string, 100) // this is how buffered channel is created
	// this will not blocking like aboe one

	// so even if we do send receive in main go routine
	// it will not stuck in dead lock

	emailChan <- "raman@gmail.com"
	emailChan <- "new@gmail.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	// if we want email worker that will take emails from queue and send email as they come
	// this can  be implemented using bufferd channel like below
	emailWorkerChan := make(chan string, 100) // buffered channel
	syncMailRoutine := make(chan bool)        // unbuffered channel for syncronization

	go emailSender(emailWorkerChan, syncMailRoutine)

	for i := range 100 {
		emailWorkerChan <- fmt.Sprintf("%d@gmail.com", i)
		//Sprintf formats according to a format specifier and returns the resulting string.
	}
	close(emailWorkerChan) // without this it will give deadlock error since that emailsender will always waits for data in this channel and defer will not run
	<-syncMailRoutine
	// this will give deadlock error after all thing will get done
	// because above in emailSender it will  always wait for email worker chan data
	// so this can we solved using closing a channel after putting all the emails

	// Topic 7
	// We can also receive data from multiple channels

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan 2", chan2Val)
		}
	}

	// Topic 8
	// We can also restrict access to channel either allow receive either sending
	// because in production we don't want even by mistakely in go routine that we
	// are using has taking channel for receiving data only
	// and we must need to assure that it will not able to add data
	// this is done as below

	// 	func emailSender(emailChan <-chan string, done chan bool) {
	// ** <-chan **
	// ** chan<- **
	//  email chan can now used to receive data only it will not allow send data to that channel
	// done chan can now used to only send data it will not allow to receive data
	// 	defer func() {
	// 		done <- true
	// 	}()
	// 	for email := range emailChan {
	// 		fmt.Println("sending email to", email)
	// 	}
	// }
}
