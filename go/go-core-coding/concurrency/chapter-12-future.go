package main

import (
	"fmt"
	"time"
)

// 一个查询结构体
// 这里的 sql 和 result 是一个简单的抽象, 具体的应用可能是更复杂的数据类型
type query struct {
	// 参数 channel
	sql chan string

	// 结果 channel
	result chan string
}

// 执行 query
func execQuery(q query) {
	// 启动 goroutine
	go func() {
		sql := <-q.sql
		// 访问数据库
		// ...

		// 输出结果
		q.result <- "result from " + sql
	}()
}

func main() {
	// 初始化 Query
	q := query{sql: make(chan string, 1), result: make(chan string, 1)}

	// 执行 Query, 注意执行的时候无需准备参数
	go execQuery(q)

	// 发送参数
	q.sql <- "SELECT * FROM <table>"

	// 做其他事情
	time.Sleep(1 * time.Second)

	// 获取结果
	fmt.Println(<-q.result)
}
