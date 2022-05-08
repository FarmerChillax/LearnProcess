package main

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	// 创建任务通道
	taskChan := make(chan task, 10)

	// 创建结果通道
	resultChan := make(chan int, 10)

	// wait 用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	// 初始化 task 的 goroutine，计算 100 个自然数之和
	go InitTask(taskChan, resultChan, 100)

	// 每个task启动一个 goroutine 处理
	go DistributeTask(taskChan, wait, resultChan)

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
func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}

	wait.Wait()
	close(result)
}

// goroutine 处理具体工作, 并将结果发送到结果通道中
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// 读取结果通道, 汇总结果
func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}
