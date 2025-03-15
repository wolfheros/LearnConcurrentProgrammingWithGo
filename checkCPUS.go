package main

import (
	"fmt"
	"runtime"
)

func main1() {
	go sayhellow()
	runtime.Gosched()
	fmt.Printf("finished!")
}

func sayhellow() {
	fmt.Printf("hello")
}
