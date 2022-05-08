package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}

		fmt.Println(sum)
		c <- struct{}{}
	}(c)

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	// 读取通道 c, 通过管道进行同步等待
	<-c
}
