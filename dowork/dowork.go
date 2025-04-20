package main

import (
	"fmt"
	"sync"
)

func doWork(cond *sync.Cond) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	cond.L.Lock()
	cond.Signal() // ALways use [Signal(), Broadcast(), Wait()]
	// when holding the mutex lock to avoid synchronization
	cond.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	for i := 0; i < 50000; i++ {
		go doWork(cond)
		fmt.Println("Waiting for child goroutine ")
		cond.Wait() // waiting for finsh signal
		fmt.Println("Chiled goroutine finished")
	}
	cond.L.Unlock()
}
