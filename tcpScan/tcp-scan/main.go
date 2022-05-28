package main

import (
	"tcp-scan/pkg"
	"time"

	"github.com/pkg/profile"
)

var TIMEOUT = 1 * time.Second

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 加载环境变量
	// ...
	// 得到A-D网段
	segmentList := []string{"10", "9", "8"}
	// 开始创建扫描任务
	// ...

}

func InitTask(endpoint string, taskChan chan<- pkg.Task, timeout ...time.Duration) {
	task := pkg.NewTask(endpoint, timeout...)
	taskChan <- *task
}

func worker(tasks, result chan pkg.Task) {
	for task := range tasks {
		if task.Error = task.Do(); task.Error == nil {
			task.Status = true
		}
		result <- task
	}
}
