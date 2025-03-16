package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	monney := 100
	go stingdy(&monney)
	go spendy(&monney)

	time.Sleep(2 * time.Second)
	fmt.Println("Monney in bank account: ", monney)
}
func stingdy(monney *int) {
	for i := 0; i < 1000000; i++ {
		*monney += 10
		runtime.Gosched()
	}
	fmt.Println("Stingdy Done!")
}

func spendy(monney *int) {
	for i := 0; i < 1000000; i++ {
		*monney -= 10
		runtime.Gosched()
	}

	fmt.Println("Spendy Done!")
}
