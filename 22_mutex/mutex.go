package main

import (
	"fmt"
	"sync"
)

// mutex = mutual exclusion
// mutex is used to protect resources from being accessed concurrently
// by multiple go routines, preventing data races and ensuring data consistency

// Purpose:

//     Mutual Exclusion:
//     A mutex ensures that only one goroutine can hold the lock and access a critical section of code or shared data at any given time.
//     Preventing Race Conditions:
//     By enforcing mutual exclusion, mutexes prevent race conditions where the outcome of an operation on shared data depends on the unpredictable order of execution of concurrent goroutines.

// for example

// mutex comes from sync package
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	// that lock will prevent other go routine access to this while we edit it see below
	// we are locking here
	p.mu.Lock()
	p.views += 1

	// we are unlocking it here
	p.mu.Unlock() // best practice is to put it in defer generally
}

func main() {
	myPost := post{views: 0}
	wg := sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		// myPost.inc() // this will syncronously update so one by one
		go myPost.inc(&wg) // this will do it asyncronously doesnt block execution
		// this is done concurently they tries modify mypost view
		// so this will met race condition multiple concurent process tries to change this data at the same time
		// to resovle this mutex was introduced
		// mutex locks resources and allows access to it to one routine at time
		// this will prevent race condition

		// how to implement this see above in struct declaration
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
