package main

import (
	"flag"
	"fmt"
	"tcp-scan/pkg"
	"time"

	"github.com/pkg/profile"
)

var TIMEOUT = 1 * time.Second

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// 加载环境变量
	// ...
	segmentA := flag.String("a", "192", "网段A")
	segmentB := flag.String("b", "%d", "网段B")
	segmentC := flag.String("c", "%d", "网段C")
	segmentD := flag.String("d", "%d", "网段D")
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	fmt.Println(*segmentA, *segmentB, *segmentC, *segmentD)
	// 得到A-D网段
	// segmentList := []string{"10", "9", "10"}
	// 开始创建扫描任务
	// ...

	// test-debug
	// segment := pkg.CreateSegment(segmentList...)
	// fmt.Println(segment)
	// // segmentChan := make(chan string, 20)
	// str := "%d.%d.%d.%d"
	// fmt.Println(strings.HasSuffix(str, "%d"))
}

func Create(segments []string) {

}

func CreateTasks(start, end int, taskChan chan<- pkg.Task, segments ...string) {

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
