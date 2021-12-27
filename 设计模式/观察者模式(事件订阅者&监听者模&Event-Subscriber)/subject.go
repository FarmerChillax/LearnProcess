package main

// 定义接口
type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}
