package main

import (
	"fmt"
	"sync"
)

// Go Routine - Go routine are light weight threads
// If we want to do multithreading or want to have concurency we use Go Routine
// Go Routine is very powerfull feature of Go

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		task(i)
// 	}
// }

// Above task func will run one by one
// one after another and will get executed in order
// this task are blocking task means task coming after needs to wait for task which is executing

// This task runs on single thread just like javascript

// If we want to run it in multiple threads parellaly than we use Go Routins
// Go Routine lets us do task in multiple thread concurently
// it will run make new light weight thread and execute whatever we have using go routine
// it will not run on actual thread but it creates light weight performant thread
// Go routine are managed by go runtime schedular
// Reference = https://medium.com/@mail2rajeevshukla/unlocking-the-power-of-goroutines-understanding-gos-lightweight-concurrency-model-3775f8e696b0

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

// func main() {
// 	for i := 0; i < 100; i++ {
// 		go task(i)
// 		// this are non blocking code so while the time this is getting executed
// 		// out main function will not have anything to execute hence
// 		// it will not print it will exit out code execution but go routin will work without blocking thread

// 	}
// 	// to block thread and see output of execution of go routine
// 	time.Sleep(time.Second * 5)
// }

// Wait Group
// since we know that go routine are non blocking
// so if main func doesn't have anything
// our program will exit and
// go routine work in background
// so if we want to see what currently is happening in go routine funcitons and block thread
// we have concept of Wait Group that will make this code syncronus and wait till go routing completes execution
// basically we make wait group block thread till it become 0
// let see

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // If we use defer keyword it whatever we have after will get executed at very end of out function just like return ()=> in js
	// wg.Done()will do nothing but decrements waitgroup task counter
	fmt.Println("executing task", id)
}

func main() {
	// this is how we create waitgroup using sync package
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1) // this will increace waitgroup task counter by one
		go task(i, &wg)
		// this are non blocking code so while the time this is getting executed
		// out main function will not have anything to execute hence
		// it will not print it will exit out code execution but go routin will work without blocking thread

	}
	wg.Wait() // this will wait until task counter will get zero
	// you can see above it will only get to zero when everything will get executed in go routine
	// this is how we block thread till go routine complete execution using wait group
}
