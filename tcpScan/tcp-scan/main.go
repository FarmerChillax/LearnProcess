package main

import (
	"tcp-scan/pkg"
	"time"

	"github.com/pkg/profile"
)

var TIMEOUT = 1 * time.Second

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	segmentList := []string{"10", "9", "10"}
	// 加载环境变量
	segment := pkg.CreateSegment(segmentList...)

	// segmentChan := make(chan string, 20)

}

func InitTask(taskChan chan<- pkg.Task) {}

func worker(tasks, result chan pkg.Task) {
	for task := range tasks {
		if task.Error = task.Do(); task.Error == nil {
			task.Status = true
		}
		result <- task
	}
}
