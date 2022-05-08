package main

import (
	"fmt"
)

const (
	NUMBER = 10
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

// 任务处理: 计算begin到end的和
// 执行结果写入到结果通道
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	workers := NUMBER

	// 工作通道
	taskChan := make(chan task, 10)

	// 结果通道
	resultChan := make(chan int, 10)

	// 信号通道
	done := make(chan struct{}, 10)

	// 初始化 task 的 goroutine，计算 100 个自然数之和
	go InitTask(taskChan, resultChan, 100)

	// 分发任务到每个 worker 中
	DistributeTask(taskChan, workers, done)

	// 获取各个 goroutine 处理完任务的通知, 并关闭结果通道
	go CloseResult(done, resultChan, workers)
	
	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultChan)
	fmt.Println(sum)
}

// 构建task并写入task通道
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tsk
	}
	close(taskChan)
}

// 读取task chan, 每个task启动一个 worker goroutine进行处理
// 并等待每个task运行完, 关闭结果通道
func DistributeTask(taskChan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskChan, done)
	}
}

// goroutine 处理具体工作, 并将结果发送到结果通道中
func ProcessTask(taskChan <-chan task, done chan struct{}) {
	for t := range taskChan {
		t.do()
	}
	done <- struct{}{}
}

// 通过 done channel 同步等待所有 worker 的结束, 然后关闭结果 channel
func CloseResult(done chan struct{}, resultChan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultChan)
}

// 读取结果通道, 汇总结果
func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}
