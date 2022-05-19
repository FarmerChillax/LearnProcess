package main

import (
	"fmt"
	"strings"
	"tcp-scan/pkg"
	"time"

	"github.com/pkg/profile"
)

var TIMEOUT = 1 * time.Second

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 读取环境变量配置
	segmentList := []string{"10", "9", "10"}
	// 加载环境变量
	segment := pkg.CreateSegment(segmentList...)
	fmt.Println(segment)
	// segmentChan := make(chan string, 20)
	str := "%d.%d.%d.%d"
	fmt.Println(strings.HasSuffix(str, "%d"))
}

func CreateTasks(segment, start, end string, taskChan chan<- pkg.Task) {}

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
