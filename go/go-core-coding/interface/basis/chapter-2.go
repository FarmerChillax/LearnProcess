package main

import "fmt"

type Printer interface {
	Print()
}

type S struct{}

func (s S) Print() {
	fmt.Println("print")
}

func main() {
	var ptr Printer

	// 没有初始化的接口调用方法会产生 panic
	// panic: runtime error: invalid memory address or nil pointer dereference
	// ptr.Print()

	// 初始化
	ptr = S{}
	ptr.Print()
}
