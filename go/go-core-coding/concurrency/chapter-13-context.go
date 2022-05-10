package main

import (
	"context"
	"fmt"
	"time"
)

// 定义一个包含 Context 的新字段
type otherContext struct {
	context.Context
}

func main() {

	// 使用 context.Background() 构建一个 WithCancel 类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())
	go worker(ctxa, "worker-1")

	// 使用 WithDeadline 包装前面的上下文对象 ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go worker(ctxb, "worker-2")

	// 使用 WithValue 包装前面的上下文对象 ctxb
	oc := otherContext{Context: ctxb}
	ctxc := context.WithValue(oc, "key", "ctx test value")

	go workWithValue(ctxc, "worker-3")

	// 故意 ”Sleep“ 10秒，让 worker2、worker3 超时退出
	time.Sleep(10 * time.Second)

	// 显式调用 worker1 的 cancel 方法通知其退出
	cancel()

	// 等待worker1打印信息
	<-ctxa.Done()
	fmt.Println("main get <-ctxa.Done()")
	time.Sleep(time.Second)
	fmt.Println("main stop")

}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get message to cancel \n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(time.Second * 1)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get message to cancel \n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s \n", name, value)
			time.Sleep(time.Second * 1)
		}
	}
}
