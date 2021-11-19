package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	mutex     sync.Mutex
	resources chan io.Closer // 将资源存放在管道中
	factory   func() (io.Closer, error)
	closed    bool
}

// 定义代理池相关的错误
var ErrPoolClosed = errors.New("pool has been closed")

// New一个代理池
// 1. 需要一个可以分配资源的函数
// 2. 规定代理池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// 从代理池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有资源剩余
	case resource, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			// 通道已关闭
			return nil, ErrPoolClosed
		}
		return resource, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// 将使用后的资源放回代理池中
// (释放占资源)
func (p *Pool) Release(r io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		// 如果代理池关闭了，就将资源关闭
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		// 尝试将该资源放入队列中
		log.Println("Release:", "In Queue")
	default:
		// 资源池满了则直接销毁这个资源
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close函数会让资源池停止工作
func (p *Pool) Close() {
	// 保证本操作与Release操作的安全
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// 如果资源池已标记为关闭，则什么也不做
	if p.closed {
		return
	}

	// 关闭资源池
	p.closed = true

	// 清空管道内的资源，将通道关闭
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
