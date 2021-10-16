package main

// 每台计算机都有如下功能
// 1. 打印
// 2. 安装打印机驱动
type computer interface {
	print()
	setPrinter(printer)
}
