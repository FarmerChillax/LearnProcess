package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			// 0 或 1 的写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
