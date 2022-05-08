package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前的 GOMAXPROCS 值
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// 设置 GOMAXPROCS 的值为 2
	runtime.GOMAXPROCS(2)

	// 获取当前的 GOMAXPROCS 值
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
