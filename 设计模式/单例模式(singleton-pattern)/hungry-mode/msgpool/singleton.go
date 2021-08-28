package msgpool

import "sync"

type Message struct {
	Count int
}

// 概述：保证全局仅有一个实例，并且提供一个全局访问点
// 注意：
// 1. 要限制调用者直接实例化该对象 -> go中利用包的访问规则->小写
// 2. 为该对象提供全局唯一的访问点
type messagePool struct {
	pool *sync.Pool
}

// 单例
// 饿汉模式: 系统加载时完成初始化
var msgPool = &messagePool{
	pool: &sync.Pool{New: func() interface{} {
		return &Message{Count: 0}
	}},
}

// 访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

// 往消息池添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

// 从消息池获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
