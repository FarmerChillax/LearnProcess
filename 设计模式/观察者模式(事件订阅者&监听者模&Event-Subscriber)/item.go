package main

import (
	"log"
	"sync"
)

// 具体项目实现
type item struct {
	observerList []observer
	name         string
	mutex        *sync.RWMutex
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

// 注册订阅对象
func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

// 取消订阅
func (i *item) deregister(o observer) {
	for index, obs := range i.observerList {
		if obs.getID() == o.getID() {
			i.mutex.Lock()
			removeObserver := obs.getID()
			i.observerList = append(i.observerList[:index], i.observerList[index+1:]...)
			i.mutex.Unlock()
			log.Printf("Remove observer: %v\n", removeObserver)
			return
		}
	}
	log.Printf("Observer at: %s not found\n", o.getID())
}

// 发布订阅
func (i *item) notifyAll() {
	for _, item := range i.observerList {
		// 更新产品名字
		item.update(i.name)
	}
}
