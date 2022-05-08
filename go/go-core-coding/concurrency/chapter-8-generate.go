package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// done 接收通知退出信号
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

// done 接收通知退出信号
func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-GenerateIntA(send):
			case ch <- <-GenerateIntB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 创建一个作为接收退出信号的 chan
	done := make(chan struct{}, 0)

	// start generate
	ch := GenerateInt(done)

	fmt.Println("Now goroutine number:", runtime.NumGoroutine())

	// 消费者消费
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 通知生产者停止生产
	done <- struct{}{}
	fmt.Println("stop generate", runtime.NumGoroutine())
}
