package main

func main() {
	// 初始化商品
	Apple := newItem("Apple iphone8")
	// 初始化两个顾客
	observerFarmer := customer{id: "farmer-chong@qq.com"}
	observerXiaotao := customer{id: "xiaotao233@qq.com"}
	// 顾客订阅产品
	Apple.register(&observerFarmer)
	Apple.register(&observerXiaotao)
	// 产品更新，通知顾客
	Apple.notifyAll()

	// 产品发生变动，更新推送
	Apple.name = "Apple iphone13"
	Apple.notifyAll()
}
