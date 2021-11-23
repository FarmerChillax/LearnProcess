package work

import "sync"

// Worker必须满足接口类型
// 才能用worker pool
type Worker interface {
	Task()
}

// Pool提供一个worker的goroutine池，这个池
// 可以完成任何已提交的worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	// 创建一个pool
	p := Pool{
		work: make(chan Worker),
	}

	// 开启worker池
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// 从通道中获取任务，并运行
			for w := range p.work {
				w.Task()
			}
			// 通道关闭后将计数器减一
			p.wg.Done()
		}()
	}

	return &p
}

// Run提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown等待的所有goroutine
func (p *Pool) Shutdown() {
	close(p.work)
	// 等待waite group结束
	p.wg.Wait()
}
