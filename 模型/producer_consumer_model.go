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
		// 往管道写元素(生产)
		product <- i
		i++
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
}

// 消费者, 参数是管道
func consumer(product <-chan int) {
	// 从管道中获取生产者生产的东西(消费)
	for p := range product {
		fmt.Printf("消费者: 消费了 %d\n", p)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	// 初始化一条传输产品的通道(管道)，buf = 10
	product := make(chan int, 10)
	// 上下文控制，这里没啥用，就是为了监听键盘的中断信号，停止程序
	ctx, cancel := context.WithCancel(context.Background())
	// 相当于finally，保证函数结束前关闭，类似python的with语法糖
	defer cancel()

	// 启动goroutine (协程or线程 看go调度器)
	go producer(product)
	go consumer(product)

	//	监听键盘的中断信号
	go func() {
		var s string
		fmt.Scanln(&s)
		cancel()
	}()

	// 当键盘中断信号发出时，这里会被唤醒(wake up)
	<-ctx.Done()
	fmt.Println("the end.")
}
