package main

import "fmt"

type canon struct {
}

func (c *canon) printFile() {
	fmt.Println("佳能打印机正在打印...")
}
