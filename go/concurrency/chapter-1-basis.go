package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println(sum)
	time.Sleep(1 * time.Second)
}

func main() {
	go sum()

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
}
