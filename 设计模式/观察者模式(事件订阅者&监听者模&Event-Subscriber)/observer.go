package main

// 观察者
type observer interface {
	update(string)
	getID() string
}
