package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
				// 增加一路监听
				// 对推出通知信号 done 的监听
			case v := <-done:
				fmt.Println(v)
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

}
