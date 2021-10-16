package main

import "fmt"

// 打印机A-> 惠普打印机
type hp struct {
}

func (h *hp) printFile() {
	fmt.Println("惠普打印机正在打印...")
}
