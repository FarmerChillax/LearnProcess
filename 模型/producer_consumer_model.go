package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 生产者
func producer(product chan<- int) {
	i := 0
	for {
		fmt.Printf("生产者: 生产了 %d\n", i)
		product <- i
		i++
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
}

// 消费者
func consumer(product <-chan int) {
	for p := range product {
		fmt.Printf("消费者: 消费了 %d\n", p)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	// 产品通道，buf = 10
	product := make(chan int, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go producer(product)
	go consumer(product)

	go func() {
		var s string
		fmt.Scanln(&s)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("the end.")
}
