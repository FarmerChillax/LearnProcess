package main

func main() {
	// 声明打印机
	hp := hp{}
	canon := canon{}

	// 声明电脑
	windows := windows{}
	mac := mac{}

	// 为windows设置惠普(hp)打印机驱动
	windows.setPrinter(&hp)
	// 为mac设置佳能(canon)打印机的驱动
	mac.setPrinter(&canon)

	// 开始打印
	windows.print()
	mac.print()

	// windows切换打印机
	windows.setPrinter(&canon)
	windows.print()
}
