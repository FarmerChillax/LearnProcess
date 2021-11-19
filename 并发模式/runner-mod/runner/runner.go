package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超时时间内执行一组任务
// 并且在操作系统发送中断任务时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统 发送信号
	interrupt chan os.Signal

	// complete 通道报告已处理完成的任务
	complete chan error

	// timeout 报告超时任务
	timeout <-chan time.Time

	// 函数, 任务列表
	tasks []func(int)
}

// 超时错误
var ErrTimeout = errors.New("received timeout")

// 接收到操作系统事件时返回
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 添加任务到Runner上。这个任务是一个
// 接受一个int类型为ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 执行tasks中的每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已注册的任务
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件北触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true

	// 继续正常运行
	default:
		return false
	}
}

// Start函数会执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务完成时发出信号
	case err := <-r.complete:
		return err
	// 当处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}

}
