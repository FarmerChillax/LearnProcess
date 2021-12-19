package day1

import (
	"log"
	"testing"
)

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func TestStackQueue(t *testing.T) {
	queue := Constructor()
	// 1
	queue.AppendTail(233)
	result := queue.DeleteHead()
	excepted := 233
	handleErr(result, excepted)
	// 2
	result = queue.DeleteHead()
	excepted = -1
	handleErr(result, excepted)
	// 3
	queue.AppendTail(123)
	queue.AppendTail(233)
	queue.AppendTail(666)
	result = queue.DeleteHead()
	excepted = 123
	handleErr(result, excepted)
}

func handleErr(result, excepted interface{}) {
	if result != excepted {
		log.Fatalf("not match, result: %v; except: %v\n", result, excepted)
	}
}
