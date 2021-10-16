package main

// 打印机接口，打印机都有如下的功能：
// 1. 打印文件
type printer interface {
	printFile()
}
