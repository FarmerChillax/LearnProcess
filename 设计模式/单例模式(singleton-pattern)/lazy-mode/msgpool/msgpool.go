package msgpool

import "sync"

type Message struct {
	Count int
}

type messagePool struct {
	pool *sync.Pool
}

var once = &sync.Once{}

// 消息池单例，首次调用时初始化（懒加载）
var msgPool *messagePool

// 获取消息池的全局唯一方法
func Instance() *messagePool {
	once.Do(func() {
		msgPool = &messagePool{
			pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
		}
	})
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
