package main

import (
	"fmt"
	// "runtime"
	"sync"
	"time"
)

func main() {
	monney := 100
	mutex := sync.Mutex{}
	go stingdy(&monney, &mutex)
	go spendy(&monney, &mutex)

	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Monney in bank account: ", monney)
	mutex.Unlock()
}
func stingdy(monney *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*monney += 10
		mutex.Unlock()
		// runtime.Gosched()
	}
	fmt.Println("Stingdy Done!")
}

func spendy(monney *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*monney -= 10
		mutex.Unlock()
		// runtime.Gosched()
	}

	fmt.Println("Spendy Done!")
}
